package handlers

import (
	"errors"
	"math/rand"
	"net/http"
	"strings"

	"github.com/cbochs/random-topics/src/models"
	"github.com/cbochs/random-topics/src/random"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// OpenSession (POST /sessions/open) Open a new random topics session
func OpenSession(c *gin.Context) {
	session := models.Session{
		Code:   random.RandomString(4),
		Status: "open",
	}
	if err := models.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": session})
}

// GetSession (GET /sessions/:code) Gets a session's information
func GetSession(c *gin.Context) {
	var session models.Session
	if err := models.DB.Where("code = ?", c.Param("code")).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": session})
}

// CloseSession (PUT /sessions/close) Close an active session
func CloseSession(c *gin.Context) {
	var input models.SessionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.
		Model(&models.Session{}).
		Where("code = ?", strings.ToUpper(input.Code)).
		Update("status", "closed")

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"data": models.SessionClosedResult{
			Code:   strings.ToUpper(input.Code),
			Status: "closed",
			OK:     false,
			Reason: "", // already closed or could not find
		}})
		return
	}

	if err := shuffleTopics(strings.ToUpper(input.Code)); err != nil {
		c.JSON(http.StatusOK, gin.H{"data": models.SessionClosedResult{
			Code:   strings.ToUpper(input.Code),
			Status: "closed",
			OK:     false,
			Reason: "No enough participants.",
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": models.SessionClosedResult{
		Code:   strings.ToUpper(input.Code),
		Status: "closed",
		OK:     true,
		Reason: "",
	}})
}

func shuffleTopics(code string) error {
	var session models.Session
	if err := models.DB.Where("code = ?", code).First(&session).Error; err != nil {
		return err
	}

	topics := []models.Topic{}
	if err := models.DB.Where("session_id = ?", session.ID).Find(&topics).Error; err != nil {
		return err
	}

	if len(topics) < 2 {
		return errors.New("Not enough topic submissions to shuffle")
	}

	picked := make([]bool, len(topics))
	for i, topic := range topics {
		for {
			assignedID := rand.Intn(len(topics))
			if assignedID != i && !picked[assignedID] {
				picked[assignedID] = true
				if err := models.DB.Model(&topic).Update("assigned", topics[assignedID].Submitted).Error; err != nil {
					return err
				}
				break
			}
		}
	}

	return nil
}

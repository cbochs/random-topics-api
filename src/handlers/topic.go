package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/cbochs/random-topics-api/src/models"
	"github.com/cbochs/random-topics-api/src/random"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SubmitTopic (POST /topics) Submit a topic to an open session
func SubmitTopic(c *gin.Context) {
	var input models.TopicInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var session models.Session
	if err := models.DB.Where("code = ?", input.Code).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if session.Status == "closed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session is closed"})
		return
	}

	topic := models.Topic{
		Code:      random.RandomString(4),
		Submitted: strings.ToLower(input.Topic),
		SessionID: session.ID,
	}
	if err := models.DB.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": topic})
}

// UpdateTopic (PUT /topics/:code) Update a topic in an open session
func UpdateTopic(c *gin.Context) {
	var input models.TopicInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var topic models.Topic
	if err := models.DB.Where("code = ?", input.Code).First(&topic).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if topic.Session.Status == "closed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session is closed"})
		return
	}

	topic.Submitted = input.Topic
	if err := models.DB.Model(&topic).Update("submitted", input.Topic).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": topic})
}

// GetSubmittedTopic (GET /topic/:code/submitted) Get a users' submitted topic
func GetSubmittedTopic(c *gin.Context) {
	var topic models.Topic
	if err := models.DB.Where("code = ?", c.Param("code")).First(&topic).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var session models.Session
	if err := models.DB.Where("id = ?", topic.SessionID).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": models.TopicResult{
		Code:        topic.Code,
		SessionCode: session.Code,
		Topic:       topic.Submitted,
		Type:        "submitted",
		OK:          true,
		Reason:      "",
	}})
}

// GetAssignedTopic (GET /topic/:code/assigned) Get a users' assigned topic
func GetAssignedTopic(c *gin.Context) {
	var topic models.Topic
	if err := models.DB.Where("code = ?", c.Param("code")).First(&topic).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var session models.Session
	if err := models.DB.Where("id = ?", topic.SessionID).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if session.Status == "open" {
		c.JSON(http.StatusOK, gin.H{"data": models.TopicResult{
			Code:        topic.Code,
			SessionCode: session.Code,
			Topic:       "",
			Type:        "assigned",
			OK:          false,
			Reason:      "Session is still open.",
		}})
		return
	} else if session.Status == "closed" && topic.Assigned == "" {
		c.JSON(http.StatusOK, gin.H{"data": models.TopicResult{
			Code:        topic.Code,
			SessionCode: session.Code,
			Topic:       "",
			Type:        "assigned",
			OK:          false,
			Reason:      "There were not enough participants.",
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": models.TopicResult{
		Code:        topic.Code,
		SessionCode: session.Code,
		Topic:       topic.Assigned,
		Type:        "assigned",
		OK:          true,
		Reason:      "",
	}})
}

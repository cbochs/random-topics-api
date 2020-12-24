package main

import (
	"math/rand"
	"time"

	"github.com/cbochs/random-topics-api/config"
	"github.com/cbochs/random-topics-api/handlers"
	"github.com/cbochs/random-topics-api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config.Get()
	models.ConnectDatabase(cfg.GetDSNStr())

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/sessions/open", handlers.OpenSession)
	r.GET("/sessions/:code", handlers.GetSession)
	r.PUT("/sessions/close", handlers.CloseSession)
	r.POST("/topics", handlers.SubmitTopic)
	r.PUT("/topics/:code", handlers.UpdateTopic)
	r.GET("/topics/:code/submitted", handlers.GetSubmittedTopic)
	r.GET("/topics/:code/assigned", handlers.GetAssignedTopic)

	r.Run()
}

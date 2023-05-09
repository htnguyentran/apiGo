package server

import (
	"fmt"
	"gin-gonic/handlers"
	"gin-gonic/rabbitmq"
	"gin-gonic/statics"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	queueName := fmt.Sprintf(statics.QueueName, os.Getenv("PREFIX_QUEUE_NAME"))
	queueRetryName := fmt.Sprintf(statics.QueueNameRetry, os.Getenv("PREFIX_QUEUE_NAME"))
	rabbitMq := rabbitmq.New(os.Getenv("AMQP_URL"), queueName, queueRetryName)
	thongnthLogictis := handlers.NewHandler(rabbitMq)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "%s", "go-hook-logistic")
	})

	router.POST("/sc/v1/shipment/:OrderNumber/status", thongnthLogictis.UpdateStatus)

	return router
}

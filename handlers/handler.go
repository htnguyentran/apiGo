package handlers

import (
	"gin-gonic/rabbitmq"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	UpdateStatus(c *gin.Context)
}

func NewHandler(rabbitMq *rabbitmq.RabbitMQ) Handler {
	return &thongnthLogictis{rabbitMQ: rabbitMq}
}

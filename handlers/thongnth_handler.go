package handlers

import (
	"encoding/json"
	"fmt"
	"gin-gonic/models"
	"gin-gonic/rabbitmq"
	"gin-gonic/statics"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type thongnthLogictis struct {
	rabbitMQ *rabbitmq.RabbitMQ
}

var timeString = time.Now().Format(time.RFC3339Nano)

func (thongnth *thongnthLogictis) UpdateStatus(c *gin.Context) {
	orderNumber := c.Param("OrderNumber")

	var model models.UpdateOrderStatusReq
	if err := c.ShouldBindJSON(&model); err != nil {
		log.Errorf("can parser req body err SC err %v", err)
		c.JSON(http.StatusOK, gin.H{"result": false, "msg": err.Error()})
		return
	}

	model.OrderNumber = orderNumber

	j, _ := json.Marshal(model)
	log.Info("receive the msg SC: ", string(j), " |time: ", timeString)
	queueName := fmt.Sprintf(statics.QueueName, os.Getenv("PREFIX_QUEUE_NAME"))
	log.Info("env : ", queueName)
	log.Info("push msg done SC", string(j), " |time: ", timeString)
	c.JSON(http.StatusOK, gin.H{"result": true, "msg": ""})
}

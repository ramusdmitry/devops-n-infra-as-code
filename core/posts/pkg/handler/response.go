package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	postApp "posts-app-service/pkg/model"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

type getAllPostsResponse struct {
	Data []postApp.Post `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, logMessage, message string) {
	logrus.Errorf("[%s] %s, cause: %s", time.Now().UTC().Format("2006-01-02 15:04:05"), logMessage, message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}

func newResponse(message string) {
	logrus.Printf("[%s] %s", time.Now().UTC().Format("2006-01-02 15:04:05"), message)
}

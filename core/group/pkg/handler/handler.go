package handler

import (
	"github.com/gin-gonic/gin"
	"group-app-service/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.Default()

	router.Use(CORS)

	api := router.Group("/api")
	{
		group := api.Group("/group")
		{
			group.GET("", h.adminIdentify, h.getUsers)
			group.PUT("", h.adminIdentify, h.updateUsers)
			group.DELETE("", h.adminIdentify, h.deleteUsers)
		}
	}
	return router
}

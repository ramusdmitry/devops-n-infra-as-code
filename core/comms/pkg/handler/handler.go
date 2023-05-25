package handler

import (
	"comms-app-service/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
		comms := api.Group("/comms")
		{

			comms.GET("", h.getAllComments)
			comms.GET("/post/:postId", h.getCommentsInPost)

			editGroup := comms.Group("", h.userIdentify)
			{
				editGroup.POST("/add/:id", h.createComment)
				editGroup.DELETE("/:commentId", h.deleteComment)
				editGroup.PUT("/:commentId", h.updateComment)
			}

		}
	}

	return router
}

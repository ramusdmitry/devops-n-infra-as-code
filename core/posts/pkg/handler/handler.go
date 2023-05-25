package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"posts-app-service/pkg/service"
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

	router := gin.New()

	router.Use(CORS)

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.GET("", h.getAllPosts)

			editGroup := posts.Group("", h.userIdentify)
			{
				editGroup.POST("", h.createPost)
				editGroup.PUT("/:postId", h.updatePost)
				editGroup.DELETE("/:postId", h.deletePost)
			}

		}
	}

	return router
}

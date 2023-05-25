package handler

import (
	"auth-app-service/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PATCH", "DELETE", "UPDATE"},
	}))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}
	}
	return router
}

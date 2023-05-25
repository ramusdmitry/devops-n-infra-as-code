package handler

import (
	authApp "auth-app-service/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input authApp.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Auth] failed to parse sign-up JSON", err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Auth] failed to create new user in DB", err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Auth] userId does not exist or failed to generate jwt-token", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"token": token,
	})

	logrus.Infof("Created a new user with id %d", id)
}

func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Auth] failed to parse sign-in JSON", err.Error())
		return
	}

	_, err := h.services.GenerateToken(input.Username, input.Username)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "[Auth] user doesn't exist", err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Auth] userId does not exist or failed to generate jwt-token", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	logrus.Infof("[Auth] Generated for @%s new token=%s", input.Username, token)

}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"group-app-service/pkg/model"
	"net/http"
)

func (h *Handler) updateUsers(c *gin.Context) {
	var input model.UpdateUsers
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Group] invalid param", err.Error())
		return
	}

	err := h.services.Administration.UpdateUsers(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Group] failed to updated users", err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("rows were updated"),
	})
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers()
	if err != nil {
		logrus.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error when select users")
	}

	type Data struct {
		Users []model.User `json:"users"`
	}

	c.JSON(http.StatusOK, Data{
		Users: users,
	})
}

func (h *Handler) deleteUsers(c *gin.Context) {
	err := h.services.Administration.DeleteUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Group] failed to updated users", err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("rows were updated"),
	})
}

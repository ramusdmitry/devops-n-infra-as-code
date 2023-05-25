package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found", "user id not found")
		return 0, errors.New("userId not found")
	}

	userIdInt, err := userId.(int)
	if !err {
		newErrorResponse(c, http.StatusInternalServerError, "user id is invalid type", "invalid type of user id")
		return 0, errors.New("invalid type of userId")
	}

	return userIdInt, nil
}

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "[Middleware] empty auth header", "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized,
			"[Middleware] invalid size of auth header", "invalid size of auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized,
			fmt.Sprintf("[Middleware] failed to parse jwt-token %s", err.Error()),
			err.Error())
		return
	}

	c.Set(userCtx, userId)
	newResponse(fmt.Sprintf("User with id=%d successfully logged in", userId))
}

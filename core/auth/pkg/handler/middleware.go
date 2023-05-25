package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

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
	logrus.Infof("User with id=%d successfully logged in", userId)
}

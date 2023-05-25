package handler

import (
	"github.com/gin-gonic/gin"
	"group-app-service/pkg/service"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	groupCtx   = "groupId"
)

func (h *Handler) adminIdentify(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "[Middleware] empty auth header", "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "[Middleware] invalid size of auth header", "invalid size of auth header")
		return
	}

	groupId, err := service.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized,
			"[Middleware] invalid group in token", err.Error())
		return
	}

	if groupId != 1 {
		newErrorResponse(c, http.StatusForbidden, "you aren`t admin", "")
		return
	}

	c.Set(groupCtx, groupId)
}

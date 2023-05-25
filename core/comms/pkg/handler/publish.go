package handler

import (
	commsApp "comms-app-service/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createComment(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input commsApp.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Comms] failed to parse comment JSON", err.Error())
		return
	}

	commentId, err := h.services.Comms.CreateComment(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Comms] failed to create new comment in DB", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      commentId,
		"userId":  userId,
		"postId":  input.PostId,
		"content": input.Content,
	})

	newResponse(fmt.Sprintf("Created a new comment with id (%d) to post (%d) by user (%d)", commentId, input.PostId, input.Id))

}

func (h *Handler) getAllComments(c *gin.Context) {
	comments, err := h.services.Comms.GetAllComments()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Comms] failed to get all comments", err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Data: comments,
	})

	newResponse("[Comms] successfully return all comments from db")
}

func (h *Handler) getCommentsInPost(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("postId"))

	comments, err := h.services.Comms.GetCommentsByPostId(postId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("[Comms] failed to get comments from post (%d)", postId), err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Data: comments,
	})

	newResponse(fmt.Sprintf("[Comms] successfully return all comments in post (%d) from db", postId))
}

func (h *Handler) deleteComment(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Comms] invalid id param", err.Error())
		return
	}

	err = h.services.DeleteCommentById(userId, commentId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("[Comms] failed to delete comment with id (%d)", commentId), err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("commentId=%d has been deleted for userId=%d", commentId, userId),
	})

	newResponse(fmt.Sprintf("[Comms] comment (%d) was successfully deleted", commentId))

}

func (h *Handler) updateComment(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Comms] invalid id param", err.Error())
		return
	}

	var input commsApp.UpdateCommentInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Comms] invalid param", err.Error())
		return
	}

	err = h.services.Comms.UpdateCommentById(userId, commentId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("[Comms] failed to update comment (%d)", commentId), err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("commentId=%d has been updated for userId=%d", commentId, userId),
	})

	newResponse(fmt.Sprintf("[Comms] comment (%d) was successfully updated", commentId))
}

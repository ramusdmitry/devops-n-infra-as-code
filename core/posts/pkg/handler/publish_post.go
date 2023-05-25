package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	postApp "posts-app-service/pkg/model"
	"strconv"
)

func (h *Handler) getAllPosts(c *gin.Context) {

	posts, err := h.services.PostsList.GetAllPosts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Posts] failed to get posts", err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: posts,
	})
}
func (h *Handler) createPost(c *gin.Context) {

	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input postApp.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Posts] failed to parse inputted post", err.Error())
		return
	}

	postId, err := h.services.PostsList.CreatePost(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Posts] failed to insert new post into db", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
		"post_id": postId,
	})

	newResponse(fmt.Sprintf("[Posts] post (%d) was successfully created", postId))

}
func (h *Handler) updatePost(c *gin.Context) {

	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Posts] invalid param", err.Error())
		return
	}

	var input postApp.UpdatePostInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Posts] invalid param", err.Error())
		return
	}

	err = h.services.PostsList.UpdatePostById(userId, postId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("[Posts] failed to update comment (%d)", postId), err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("postId=%d has been updated for userId=%d", postId, userId),
	})

	newResponse(fmt.Sprintf("[Posts] post (%d) was successfully updated", postId))

}
func (h *Handler) deletePost(c *gin.Context) {

	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Posts] invalid id param", err.Error())
		return
	}

	err = h.services.PostsList.DeletePostById(userId, postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("[Posts] failed to delete post with id (%d)", postId), err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("postId=%d has been deleted for userId=%d", postId, userId),
	})

	newResponse(fmt.Sprintf("[Posts] post (%d) was successfully deleted", postId))

}

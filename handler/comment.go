package handler

import (
	"elearning-api/comment"
	"elearning-api/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService comment.Service
}

func NewCommentHandler(commentService comment.Service) *commentHandler {
	return &commentHandler{commentService}
}

func (h *commentHandler) CreateNewComment(c *gin.Context) {
	var input comment.CommentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Create new campaign is failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.commentService.CreateComment(input)
	if err != nil {
		response := helper.APIResponse("Failed to comment", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Comment successfuly created", "success", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) GetCommentsByTaskID(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Query("task_id"))

	fmt.Println(taskID)

	comments, err := h.commentService.FindCommentsByTaskID(taskID)
	if err != nil {
		response := helper.APIResponse("Failed to get comments", "error", http.StatusBadRequest, comments)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	if len(comments) == 0 {
		response := helper.APIResponse("Nothing comments by that taskID", "no data", http.StatusResetContent, comments)
		c.JSON(http.StatusResetContent, response)

		return
	}

	formatComment := comment.FormatComments(comments)

	response := helper.APIResponse("Successfuly get comments by that taskID", "success", http.StatusOK, formatComment)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) GetCommentsByForumID(c *gin.Context) {
	forumID, _ := strconv.Atoi(c.Query("forum_id"))

	fmt.Println(forumID)

	comments, err := h.commentService.FindCommentsByForumID(forumID)
	if err != nil {
		response := helper.APIResponse("Failed to get comments", "error", http.StatusBadRequest, comments)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	if len(comments) == 0 {
		response := helper.APIResponse("Nothing comments by that forumID", "no data", http.StatusResetContent, comments)
		c.JSON(http.StatusResetContent, response)

		return
	}

	formatComment := comment.FormatComments(comments)

	response := helper.APIResponse("Successfuly get comments by that forumID", "success", http.StatusOK, formatComment)
	c.JSON(http.StatusOK, response)
}

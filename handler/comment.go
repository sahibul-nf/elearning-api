package handler

import (
	"elearning-api/comment"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService comment.Service
}

func NewcommentHandler(s comment.Service) *commentHandler {
	return &commentHandler{s}
}

func (h *commentHandler) CreateComment(c *gin.Context) {

}

func (h *commentHandler) GetComments(c *gin.Context) {

}

func (h *commentHandler) DeleteComment(c *gin.Context) {

}

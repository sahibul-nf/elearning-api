package comment

type CommentInput struct {
	Comment string `json:"comment" binding:"required"`
}

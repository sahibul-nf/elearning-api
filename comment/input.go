package comment

type CommentInput struct {
	Comment string `json:"comment" binding:"required"`
}

type DeleteCommentInput struct {
	ID int `uri:"id" binding:"required"`
}

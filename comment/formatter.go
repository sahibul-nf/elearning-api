package comment

import "time"

type CommentFormatter struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ForumID   int       `json:"forum_id"`
	TaskID    int       `json:"task_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatComment(comment Comment) CommentFormatter {
	formatter := CommentFormatter{}
	formatter.ID = comment.ID
	formatter.UserID = comment.UserID
	formatter.ForumID = comment.ForumID
	formatter.TaskID = comment.TaskID
	formatter.Comment = comment.Comment
	formatter.CreatedAt = comment.CreatedAt

	return formatter
}

func FormatComments(comments []Comment) []CommentFormatter {
	if len(comments) == 0 {
		return []CommentFormatter{}
	}

	var commentsFormatter []CommentFormatter

	for _, comment := range comments {
		commentFormatter := FormatComment(comment)
		commentsFormatter = append(commentsFormatter, commentFormatter)
	}

	return commentsFormatter

}

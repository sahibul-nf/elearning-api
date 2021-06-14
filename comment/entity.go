package comment

import "time"

type Comment struct {
	ID        int
	UserID    int
	ForumID   int
	TaskID    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

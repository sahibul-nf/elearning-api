package comment

import "time"

type Comment struct {
	ID        int
	UserID    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

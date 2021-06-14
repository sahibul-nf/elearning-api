package comment

type Service interface {
	CreateComment(input CommentInput) (Comment, error)
	FindCommentsByTaskID(taskID int) ([]Comment, error)
	FindCommentsByForumID(forumID int) ([]Comment, error)
	DeleteComment(ID int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) CreateComment(input CommentInput) (Comment, error) {
	comment := Comment{}
	comment.Comment = input.Comment

	newComment, err := s.repository.Save(comment)

	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

func (s *service) FindCommentsByTaskID(taskID int) ([]Comment, error) {

	comments, err := s.repository.FindByTaskID(taskID)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *service) FindCommentsByForumID(forumID int) ([]Comment, error) {

	comments, err := s.repository.FindByForumID(forumID)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *service) DeleteComment(ID int) (bool, error) {

	comment, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}

	_, err = s.repository.Delete(comment.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

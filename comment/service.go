package comment

type Service interface {
	UserComment(comment CommentInput) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) UserComment(input CommentInput) (Comment, error) {
	userComment := Comment{}

	userComment.Comment = input.Comment

	comment, err := s.repository.Save(userComment)

	if err != nil {
		return comment, err
	}

	return comment, nil
}

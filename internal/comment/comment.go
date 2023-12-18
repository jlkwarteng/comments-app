package comment

import (
	"context"
	"fmt"

	"errors"
)

// Comment - A representation of the comment structure for our service
type Comment struct {
	Id     int
	Slug   string
	Body   string
	Author string
}

var (
	ErrorFetchingComment = errors.New("There was an error getting the error by ID")
	ErrorNotImplemented  = errors.New("Not Implemented")
)

// Service - A presentation of our Service Object on which all our logic will be built
type Service struct {
	Store Store
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	createComment(context.Context, Comment) (Comment, error)
	updateComment(context.Context, Comment) error
	deleteComment(context.Context, string) error
}

// NewService - is used to instantiate a new Service by returning a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) getComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment ")
	comm, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err.Error())
		return Comment{}, ErrorFetchingComment
	}
	return comm, nil
}

func (s *Service) updateComment(ctx context.Context, cmt Comment) error {
	return ErrorNotImplemented
}

func (s *Service) createComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}

func (s *Service) deleteComment(ctx context.Context, id string) error {
	return ErrorNotImplemented
}

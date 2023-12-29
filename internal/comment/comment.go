package comment

import (
	"context"
	"fmt"

	"errors"
)

// Comment - A representation of the comment structure for our service
type Comment struct {
	Id     string
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
	CreateComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) error
	DeleteComment(context.Context, string) error
}

// NewService - is used to instantiate a new Service by returning a pointer to a new service
func   NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment ")
	comm, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err.Error())
		return Comment{}, ErrorFetchingComment
	}
	return comm, nil
}

func (s *Service) UpdateComment(ctx context.Context, Id string, cmt Comment) (Comment, error) {
	err := s.Store.UpdateComment(ctx, Id, cmt)
	if err != nil {
		return Comment{}, fmt.Errorf("Failed to update Comment", err)
	}

	return Comment{}, nil

}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	comment, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, fmt.Errorf("Error Creating Comment %w", err)
	}
	return comment, nil

}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		return fmt.Errorf("Failed To Delete Comment ")
	}
	return nil

}

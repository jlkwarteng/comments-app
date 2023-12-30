//go:build integration
// +build integration

package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/jlkwarteng/comments-app/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		fmt.Println("Testing the creation of Comments ")
		cmt, err := db.CreateComment(context.Background(), comment.Comment{Slug: "Slug", Author: "Jeremiah", Body: "MyBody"})
		fmt.Sprintf("THIS IS COMMENT %v", cmt)
		newCmt, err := db.GetComment(context.Background(), cmt.Id)
		assert.Equal(t, "Slug", newCmt.Slug)
		assert.NoError(t, err)
	})

	t.Run("Test Comment Deletion", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "Test Slug",
			Author: "Test Author",
			Body:   "My Test Body",
		})

		assert.NotNil(t, cmt.Id)

		err = db.DeleteComment(context.Background(), cmt.Id)
		assert.NoError(t, err)
		_, err = db.GetComment(context.Background(), cmt.Id)

		assert.Error(t, err)
	})

	t.Run("Test Update Comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		orig_cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "My Slug",
			Author: "My Author",
			Body:   "My Body",
		})

		assert.NoError(t, err)
		assert.NotNil(t, orig_cmt.Id)

		err = db.UpdateComment(context.Background(), orig_cmt.Id, comment.Comment{
			Slug: "My New Slug",
		})

		cmt, err := db.GetComment(context.Background(), orig_cmt.Id)
		assert.Equal(t, "My New Slug", cmt.Slug)
		assert.NoError(t, err)
		// assert.Equal(t, "My New Slug", cmt.Slug)

	})
}

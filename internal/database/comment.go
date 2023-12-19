package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jlkwarteng/comments-app/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(cmtRow CommentRow) comment.Comment {
	return comment.Comment{
		Id:     cmtRow.ID,
		Slug:   cmtRow.Slug.String,
		Body:   cmtRow.Body.String,
		Author: cmtRow.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(ctx, `Select id, body, author, slug from comments where id = $1`, uuid)
	err := row.Scan(&cmtRow.ID, &cmtRow.Body, &cmtRow.Author, &cmtRow.Slug)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("Error Fetching comment by uuid %w", err)
	}
	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) CreateComment(ctx context.Context, comm comment.Comment) (comment.Comment, error) {
	comm.Id = uuid.NewV4().String()

	postRow := CommentRow{
		ID:     comm.Id,
		Slug:   sql.NullString{comm.Slug, true},
		Author: sql.NullString{comm.Author, true},
		Body:   sql.NullString{comm.Body, true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`Insert into comments values (:id, :slug, :author, :body)`, postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("Failed to Insert Comment %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("Failed to close rows %w ", err)
	}

	return comm, nil
}
func (d *Database) UpdateComment(ctx context.Context, id string, comm comment.Comment) error {

	updateRow := CommentRow{
		Slug:   sql.NullString{comm.Slug, true},
		Body:   sql.NullString{comm.Body, true},
		Author: sql.NullString{comm.Author, true},
	}

	rows, err := d.Client.NamedQueryContext(ctx, `UPDATE comment comment set body = :id, slug= :slug, author= :author where id = :id`, updateRow)
	if err != nil {
		return fmt.Errorf("Error in update Comment %w", err)
	}

	if err := rows.Close(); err != nil {
		return fmt.Errorf("Failed to close the row %w", err)
	}

	return nil
}
func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(ctx,
		`Delete from comments where id = $1`, id)
	if err != nil {
		return fmt.Errorf("Failed to Delete Comment %w", err)
	}

	return nil
}

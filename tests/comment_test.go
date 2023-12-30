//go:build e2e
// +build e2e

package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostComment(t *testing.T) {
	t.Run("Can post A comment", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().SetBody(`{"slug": "/", "body": "Body Test HEre", "author": "Jeremiah"}`).Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("Cannot Post comment without JWT", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().
			SetHeader("Authorization", "bearer 23423423423423423423423").
			SetBody(`{"slug": "/", "body": "Body Test HEre", "author": "Jeremiah"}`).Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})
}

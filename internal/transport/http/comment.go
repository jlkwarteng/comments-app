package http

import (
	"context"
	"encoding/json"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jlkwarteng/comments-app/internal/comment"
)

type CommentService interface {
	CreateComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}
type Response struct {
	Message string
}

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}
	cmt, err := h.Service.CreateComment(r.Context(), cmt)
	if err != nil {
		log.Print(err)
		return
	}
}
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}
	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfullly Deleted the Comment"}); err != nil {
		panic(err)
	}
}

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"

	"loment/models"
	"loment/repositories"
)

type contextKey string

var (
	idKey contextKey = "id"
)

// CommentHandler core handler
type CommentHandler struct {
	Repo *repositories.CommentRepository
}

// Create create comment, return id
func (handler *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	cmt := new(models.Comment)
	err := json.NewDecoder(r.Body).Decode(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	if cmt.Id == "" {
		cmt.Id = uuid.NewV4().String()
	}
	err = handler.Repo.Create(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	respondOkWithString(w, cmt.Id)
}

// Get get comment
func (handler *CommentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	result, err := handler.Repo.Get(id)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	respondOkWithObject(w, result)
}

// Delete delete comment, return id
func (handler *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	_, err := handler.Repo.Delete(id)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	respondOkWithObject(w, true)
}

// Update update comment, return id
func (handler *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	cmt := new(models.Comment)
	err := json.NewDecoder(r.Body).Decode(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	cmt.Id = id
	err = handler.Repo.Update(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	respondOkWithObject(w, true)
}

// Count comments return number
func (handler *CommentHandler) Count(w http.ResponseWriter, r *http.Request) {
	cmt := new(models.CommentQuery)
	err := json.NewDecoder(r.Body).Decode(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	total, err := handler.Repo.Count(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	respondOkWithObject(w, total)
}

// Query query comments return list of comment
func (handler *CommentHandler) Query(w http.ResponseWriter, r *http.Request) {
	cmt := new(models.CommentQuery)
	err := json.NewDecoder(r.Body).Decode(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	cmts, err := handler.Repo.Query(cmt)
	if err != nil {
		respondErrWithError(w, err)
		return
	}
	respondOkWithObject(w, cmts)
}

// ParamId parse id from url
func ParamID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), idKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func respondErrWithError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("%v", err)))
}

func respondOkWithString(w http.ResponseWriter, str string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}

func respondOkWithObject(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

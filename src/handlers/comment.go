package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type contextKey string

var (
	idKey contextKey = "id"
)

// All get all comment
func All(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("id:%s", "abc")))
}

// Create create comment
func Create(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	w.Write([]byte(fmt.Sprintf("id:%s", id)))
}

// Get get comment
func Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	w.Write([]byte(fmt.Sprintf("id:%s", id)))
}

// Delete delete comment
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	w.Write([]byte(fmt.Sprintf("id:%s", id)))
}

// Update update comment
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(idKey).(string)
	w.Write([]byte(fmt.Sprintf("id:%s", id)))
}

// Query query comments
func Query(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("id:%s", "abc")))
}

// ParamID parse id from url
func ParamID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), idKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

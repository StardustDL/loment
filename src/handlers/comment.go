package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"loment/models"
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
	result := models.Comment{
		ID: id,
	}
	respondOkWithObject(w, result)
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
	result := [5]models.Comment{}
	respondOkWithObject(w, result)
}

// ParamID parse id from url
func ParamID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), idKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func respondOkWithObject(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

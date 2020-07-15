package handlers

import (
	"fmt"
	"net/http"
)

func All(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "abc")))
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "abc")))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "abc")))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "abc")))
}

func Query(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "abc")))
}

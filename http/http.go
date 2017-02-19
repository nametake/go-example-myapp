package http

import (
	"net/http"

	"github.com/nametake/go-example-myapp"
)

func HandlerFunc(pattern string, handler http.HandlerFunc) {
	http.HandleFunc(pattern, handler)
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

type Handler struct {
	UserService myapp.UserService
}

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Users() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

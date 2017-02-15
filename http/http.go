package http

import (
	"net/http"

	"github.com/nametake/go-example-myapp"
)

func Handle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

type UserHandler struct {
	UserService myapp.UserService
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

type UsersHandler struct {
	UserService myapp.UserService
}

func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

type CreateUserHandler struct {
	UserService myapp.UserService
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

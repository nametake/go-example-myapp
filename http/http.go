package http

import (
	"net/http"

	"github.com/nametake/go-example-myapp"
)

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
	CreateUserService myapp.UserService
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

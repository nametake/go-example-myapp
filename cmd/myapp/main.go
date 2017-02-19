package main

import (
	"github.com/nametake/go-example-myapp/http"
	"github.com/nametake/go-example-myapp/mysql"
)

func main() {
	service := &mysql.UserService{}

	handler := http.Handler{}

	http.ListenAndServe(":8080", nil)
}

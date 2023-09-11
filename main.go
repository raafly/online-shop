package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializedEvent()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
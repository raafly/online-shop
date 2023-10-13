package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type CustomerController interface {
	Register(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

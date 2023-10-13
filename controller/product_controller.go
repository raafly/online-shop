package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) 
	GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) 
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) 
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
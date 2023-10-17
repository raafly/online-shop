package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request, Param httprouter.Params)
}
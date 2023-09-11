package exception

import (
	"net/http"

	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/web"
)

func ErrorHandle(w http.ResponseWriter, r *http.Request, err interface{}) {

	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r, err) 

}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok:= err.(NotFoundError)	
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: nil,
			Message: exception.Error,
		}

		helper.WriteToRequestBody(w, webResponse)
		
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
		Message: "SERVER SEDANG SIBUK",
	}

	helper.WriteToRequestBody(w, webResponse)
}
package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/web"
	"github.com/raafly/catering/service"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) *CustomerControllerImpl {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	registerCreateRequest := web.CustomerRegisterRequest{}
	helper.ReadFromRequestBody(r, &registerCreateRequest)

	customerResponse := controller.CustomerService.Register(r.Context(), registerCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "SUCCESS",
		Data: customerResponse,
		Message: "berhasil membuat account",
	}    

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *CustomerControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	loginCreateRequest := web.CustomerLoginRequest{}
	helper.ReadFromRequestBody(r, &loginCreateRequest)

	customerResponse := controller.CustomerService.Login(r.Context(), loginCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "SUCCESS",
		Data: customerResponse,
		Message: "berhasil masuk ke account",
	}

	helper.WriteToRequestBody(w, webResponse)
}


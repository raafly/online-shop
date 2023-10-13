package controller

import (
	"net/http"
	"strconv"

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

	customerResponse, token := controller.CustomerService.Login(r.Context(), loginCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "SUCCESS",
		Data: customerResponse,
		Message: "berhasil masuk ke account",
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *CustomerControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customerUpdateRequest := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(r, &customerUpdateRequest)

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	_ = controller.CustomerService.Update(r.Context(), customerUpdateRequest)
	WebResponse := web.WebResponse{
		Code: 201,
		Status: "SUCCESS",
		Message: "SUCCESS UPDATE ACCOUNT",
	}

	helper.WriteToRequestBody(w, WebResponse)
}

func (controller *CustomerControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("customerId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(r.Context(), id)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, webResponse)	
}

func (controller *CustomerControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customer := controller.CustomerService.FindById(r.Context(), id)	
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: customer,
		Message: "SUCCESS",
	}

	helper.WriteToRequestBody(w, webResponse)
}
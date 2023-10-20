package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/web"
	"github.com/raafly/catering/service"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(w http.ResponseWriter, r *http.Request, Param httprouter.Params) {
	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(r, &orderCreateRequest)

	orderResponse := controller.OrderService.Create(r.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "SUCCESS",
		Data: orderResponse,
		Message: "YOUR ORDER SUCCESS CREATED",
	} 

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *OrderControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ordertId := params.ByName("orderId")
	id, err := strconv.Atoi(ordertId)
	helper.PanicIfError(err)

	log.Println("repository controller", id)

	order := controller.OrderService.GetById(r.Context(), id)
	webResposen := web.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: order,
	}

	helper.WriteToRequestBody(w, webResposen)
}

func (controller *OrderControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	orderCreateRequest := web.OrderUpateRequest{}
	helper.ReadFromRequestBody(r, &orderCreateRequest)

	ordertId := params.ByName("orderId")
	id, err := strconv.Atoi(ordertId)
	helper.PanicIfError(err)

	orderCreateRequest.Id_order = id

	controller.OrderService.Update(r.Context(), orderCreateRequest)
	webRespose := web.WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, webRespose)
}


func (controller *OrderControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ordertId := params.ByName("orderId")
	id, err := strconv.Atoi(ordertId)
	helper.PanicIfError(err)

	controller.OrderService.Delete(r.Context(), id)
	webRespose := web.WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, webRespose)
}
package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/web"
	"github.com/raafly/catering/service"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	product := controller.ProductService.GetAll(r.Context())
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductsRequest{}
	helper.ReadFromRequestBody(r, &productCreateRequest)

	product := controller.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	product := controller.ProductService.GetById(r.Context(), id)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
	}

	helper.WriteToRequestBody(w, webResponse)

}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(r, &productUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	product := controller.ProductService.Update(r.Context(), productUpdateRequest)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
		Message: "SUCCES UPDATE",
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, webResponse)
}

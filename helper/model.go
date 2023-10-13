package helper

import (
	"github.com/raafly/catering/model/domain"
	"github.com/raafly/catering/model/web"
)

func ToCustomerResponse(customer domain.Customers) web.RegisterSuccess {
	return web.RegisterSuccess{
		Username: customer.Username,
		Email: customer.Email,
		Address: customer.Address,
		No_telp: customer.No_telp,
	}
}

func ToCustomerResponseLogin(customer domain.Customers) web.LoginSuccess {
	return web.LoginSuccess{
		Email: customer.Email,
		Role: customer.Role,
	}
}

func ToProductReponse(product domain.Products) web.ProductResponse {
	return web.ProductResponse{
		Id: product.Id,
		Name: product.Name,
		Description: product.Description,
		Quantity: product.Quantity,
		Price: product.Price,
		Category: product.Category,
	}
}

func ToProductResponses(categories []domain.Products) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range categories {
		productResponses = append(productResponses, ToProductReponse(product))
	}
	return productResponses
}


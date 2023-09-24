package web

type ProductsRequest struct {
	Name		 string		`json:"name" validate:"required"`
	Description	 string		`json:"description" validate:"required"`
	Quantity	 int		`json:"quantity" validate:"required"`
	Price		 float64	`json:"price" validate:"required"`
}
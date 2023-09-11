package web

type ProductsRequest struct {
	Id 			 int 		`json:"id"`	
	Name		 string		`json:"name" validate:"required"`
	Description	 string		`json:"description" validate:"required"`
	Quantity	 int		`json:"quantity" validate:"required"`
	Price		 int		`json:"price" validate:"required"`
}
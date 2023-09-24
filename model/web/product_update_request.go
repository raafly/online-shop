package web

type ProductUpdateRequest struct {
	Id			int		`json:"id" validate:"required"`
	Price		float64	`json:"price" validate:"required"`
	Quantity	int		`json:"quantity" validate:"required"`
}
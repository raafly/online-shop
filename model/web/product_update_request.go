package web

type ProductUpdateRequest struct {
	Id			int		`json:"id" validate:"required"`
	Quantity	int		`json:"quantity" validate:"required"`
}
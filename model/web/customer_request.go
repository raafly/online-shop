package web

type CustomerUpdateRequest struct {
	Id			int		`json:"id" validate:"required"`
	Address		string	`json:"address" validate:"required"`
}
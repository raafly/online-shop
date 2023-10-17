package web

type OrderCreateRequest struct {
	Id_order		int		`json:"id_order" validate:"required"`
	Id_user			int		`json:"id_user" validate:"required"`
	Id_product		int		`json:"id_product" validate:"required"`
	Note			string	`json:"note" validate:"ascii"`
}

type OrderSuccess struct {
	Id_product		int		`json:"id_product"`
	Note			string	`json:"note"`
}
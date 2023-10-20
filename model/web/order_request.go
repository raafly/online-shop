package web

type OrderCreateRequest struct {
	Id_order		int		`json:"id_order" validate:"required"`
	Id_user			int		`json:"id_user" validate:"required"`
	Id_product		int		`json:"id_product" validate:"required"`
	Note			string	`json:"note" validate:"ascii"`
}

type OrderDetailCreateRequest struct {
	// this struct can be use as order_detail responsee
	Id_order		int		`json:"id_order"`
	Id_product		int		`json:"id_product"`
	Id_user 		int		`json:"id_user"`
	Address			string	`json:"address"`
	No_telp			string	`json:"no_telp"`
	Name			string	`json:"name"`
	Price			uint	`json:"price"`
	Quantity		int		`json:"quantity"`
	Status			string	`json:"status"`
}

type OrderSuccess struct {
	Id_product		int		`json:"id_product"`
	Note			string	`json:"note"`
}

type OrderDetailResponse struct {
	Id_user				int		`json:"idUser"`
	Id_product			int		`json:"idProduct"`
	Id_order			int		`json:"idOrder"`
	Product_name		string	`json:"name"`
	Product_quantity	int		`json:"quantity"`
	Product_price		uint	`json:"price"`
	Customer_address	string	`json:"address"`
	Customer_noTelp		string	`json:"noTelp"`
	Order_status		string	`json:"status"`
}
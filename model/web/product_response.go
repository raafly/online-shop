package web

type ProductResponse struct {
	Name		 string		`json:"name"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity"`
	Price		 int		`json:"price"`
}
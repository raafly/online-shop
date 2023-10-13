package web

type ProductResponse struct {
	Id			 int		`json:"id"`
	Name		 string		`json:"name"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity"`
	Price		 float64	`json:"price"`
	Category	 string		`json:"category"`
}
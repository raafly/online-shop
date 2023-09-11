package domain

type Products struct {
	Id			 int		`json:"id" validate:"required"`
	Name		 string	    `json:"name"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity"`
	Price		 int		`json:"price"`
}
package web

type RegisterSuccess struct {
	Username	string	`json:"username"`  	
	Email 		string	`json:"email"`
}

type LoginSuccess struct {
	Email 		string	`json:"email"`
	Role 		string  `json:"role"`
}

type ProductCreate struct {
	Id			 int		`json:"id"`
	Name		 string	    `json:"name"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity"`
	Price		 int		`json:"price"`
}
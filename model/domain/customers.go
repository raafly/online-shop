package domain

type Customers struct {
	Id 		 int		`json:"id"`
	Username string		`json:"username"`
	Email	 string		`json:"email"`
	Password string		`json:"password"`
	Role 	 string		`json:"role"`
}
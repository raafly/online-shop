package web

type CustomerRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email	 string	`json:"email" validate:"required"`
	Password string	`json:"password" validate:"required"`
}
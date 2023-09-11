package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/catering/controller"
	"github.com/raafly/catering/exception"
)

func NewRouter(categoryController controller.CustomerController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", categoryController.Register)
	router.POST("/api/login", categoryController.Login)

	router.PanicHandler = exception.ErrorHandle

	return router
}
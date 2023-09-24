package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/catering/controller"
	"github.com/raafly/catering/exception"
)

func NewRouter(customerController controller.CustomerController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", customerController.Register)
	router.POST("/api/login", customerController.Login)

	router.GET("/api/products", productController.GetAll)
	router.POST("/api/products", productController.Create)
	router.GET("/api/products/:productId", productController.GetById)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandle

	return router
}
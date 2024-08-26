package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krnveersharma/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controllers.signUp())
	incomingRoutes.POST("user/login", controllers.Login())
	incomingRoutes.POST("admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("users/productview", controllers.SearchProduct())
	incomingRoutes.GET("users/search", controllers.SearchProductByQuery())
}

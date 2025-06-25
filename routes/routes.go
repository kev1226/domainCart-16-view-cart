package routes

import (
	"get-cart/view"

	"github.com/gin-gonic/gin"
	"github.com/kev1226/auth-common-go/jwt"
)

func RegisterCartRoutes(router *gin.Engine) {
	cart := router.Group("/cart")
	{
		cart.GET("", jwt.AuthGuard("user"), view.GetCart)
	}
}

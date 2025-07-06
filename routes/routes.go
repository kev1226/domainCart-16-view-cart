package routes

import (
	"get-cart/view"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kev1226/auth-common-go/jwt"
)

func RegisterCartRoutes(router *gin.Engine) {
	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cart := router.Group("/cart")
	{
		cart.GET("", jwt.AuthGuard("user"), view.GetCart)
	}
}

package view

import (
	"fmt"
	"get-cart/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	// Extraemos el userID y email del contexto (establecidos por AuthGuard)
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UserID not found in context"})
		return
	}

	// Convertimos userID (int) a string para uso con Redis
	userID := fmt.Sprintf("%v", userIDValue)

	// Obtener los productos del carrito
	items, err := presenter.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving cart items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

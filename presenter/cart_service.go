package presenter

import (
	"get-cart/dto"
	"get-cart/model"
)

func GetCart(userID string) ([]dto.CartItemResponse, error) {
	entities, err := model.GetCartItems(userID)
	if err != nil {
		return nil, err
	}

	var response []dto.CartItemResponse
	for _, e := range entities {
		response = append(response, dto.CartItemResponse{
			ProductID: e.ProductID,
			Name:      e.Name,
			Price:     e.Price,
			Quantity:  e.Quantity,
		})
	}

	return response, nil
}

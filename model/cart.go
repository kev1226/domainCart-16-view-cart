package model

import (
	"encoding/json"
	"fmt"
	"get-cart/config"
	"get-cart/entity"
)

func GetCartItems(userID string) ([]entity.CartItem, error) {
	key := fmt.Sprintf("cart:%s", userID)
	values, err := config.RedisClient.HVals(config.Ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var items []entity.CartItem
	for _, val := range values {
		var item entity.CartItem
		if err := json.Unmarshal([]byte(val), &item); err == nil {
			items = append(items, item)
		}
	}

	return items, nil
}

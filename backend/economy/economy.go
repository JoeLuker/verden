// app/economy/economy.go
package economy

import (
	"fmt"

	"github.com/JoeLuker/verden/models"
)

func BuyItem(buyer *models.Player, seller *models.Player, itemID string, quantity int, items map[string]models.Item) error {
	item, exists := items[itemID]
	if !exists {
		return fmt.Errorf("item not found")
	}

	totalPrice := item.Value * quantity
	if buyer.GoldCoins < totalPrice {
		return fmt.Errorf("insufficient funds")
	}
	if seller.Inventory[itemID] < quantity {
		return fmt.Errorf("seller does not have enough of the item")
	}

	buyer.GoldCoins -= totalPrice
	seller.GoldCoins += totalPrice
	seller.Inventory[itemID] -= quantity
	buyer.Inventory[itemID] += quantity

	return nil
}

func SellItem(seller *models.Player, buyer *models.Player, itemID string, quantity int, items map[string]models.Item) error {
	return BuyItem(buyer, seller, itemID, quantity, items) // Selling is just buying in reverse
}

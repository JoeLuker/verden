// app/main.go
package main

import (
	"context"
	"fmt"
	"github.com/JoeLuker/verden/app/db"
	"github.com/JoeLuker/verden/app/economy"
	"github.com/JoeLuker/verden/app/models"
)

func main() {
	mongoClient := db.ConnectMongoDB()
	defer mongoClient.Disconnect(context.Background())

	redisClient := db.ConnectRedis()
	defer redisClient.Close()

	// Sample items
	items := map[string]models.Item{
		"item1": {ID: "item1", Name: "Sword", Value: 100},
		"item2": {ID: "item2", Name: "Shield", Value: 80},
	}

	// Sample players
	player1 := &models.Player{ID: "player1", Name: "Alice", GoldCoins: 200, Inventory: make(map[string]int)}
	player2 := &models.Player{ID: "player2", Name: "Bob", GoldCoins: 150, Inventory: map[string]int{"item1": 1}}

	// Simulate a transaction
	err := economy.BuyItem(player1, player2, "item1", 1, items)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction successful")
		fmt.Printf("Player1: %+v\n", player1)
		fmt.Printf("Player2: %+v\n", player2)
	}
}

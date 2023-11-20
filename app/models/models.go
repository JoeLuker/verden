// app/models/models.go
package models

type Player struct {
	ID        string
	Name      string
	GoldCoins int
	Inventory map[string]int // Map of item ID to quantity
}

type Item struct {
	ID    string
	Name  string
	Value int // Value in gold coins
}

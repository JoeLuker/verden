package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JoeLuker/verden/models"
)

// FormStructureHandler creates a form for inputting simulation parameters.
func FormStructureHandler(w http.ResponseWriter, r *http.Request) {
	form := models.FormStructure{
		Fields: []models.FormField{
			{
				Name:         "ResourceAvailability",
				Type:         "number",
				DefaultValue: 100,
				Placeholder:  "Enter Resource Availability (0-100)",
			},
			{
				Name:         "MarketDemand",
				Type:         "number",
				DefaultValue: 50,
				Placeholder:  "Enter Market Demand (0-100)",
			},
			{
				Name:         "PoliticalStability",
				Type:         "number",
				DefaultValue: 70,
				Placeholder:  "Enter Political Stability (0-100)",
			},
			{
				Name:         "TradeRelations",
				Type:         "number",
				DefaultValue: 60,
				Placeholder:  "Enter Trade Relations (0-100)",
			},
			{
				Name:         "CraftingAbility",
				Type:         "number",
				DefaultValue: 80,
				Placeholder:  "Enter Crafting Ability (0-100)",
			},
			{
				Name:         "Negotiation",
				Type:         "number",
				DefaultValue: 60,
				Placeholder:  "Enter Negotiation Skill (0-100)",
			},
			{
				Name:         "CombatAbility",
				Type:         "number",
				DefaultValue: 50,
				Placeholder:  "Enter Combat Ability (0-100)",
			},
			{
				Name:         "MagicUse",
				Type:         "number",
				DefaultValue: 40,
				Placeholder:  "Enter Magic Use (0-100)",
			},
			{
				Name:         "Diplomacy",
				Type:         "number",
				DefaultValue: 70,
				Placeholder:  "Enter Diplomacy (0-100)",
			},
			// Add any other fields as required
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(form)
}

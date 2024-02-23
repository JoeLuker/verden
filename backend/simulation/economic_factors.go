// File: simulation/economic_factors.go
package simulation

import (
	"math"
	"math/rand"
)

type EconomicFactors struct {
	ResourceAvailability float64
	MarketDemand         float64
	PoliticalStability   float64
	TradeRelations       float64 // New factor
	// Add other factors as needed
}

func (e *EconomicFactors) Update() {
	// Dynamic update logic based on various factors
	e.ResourceAvailability = updateResourceAvailability(e.ResourceAvailability, e.PoliticalStability)
	e.MarketDemand = updateMarketDemand(e.MarketDemand, e.TradeRelations)
	e.PoliticalStability += (rand.Float64() * 10) - 5 // Random fluctuation
	e.TradeRelations += (rand.Float64() * 10) - 5     // Random fluctuation
}

func updateResourceAvailability(currentValue float64, politicalStability float64) float64 {
	// Logic to update resource availability based on political stability
	return math.Max(0, currentValue+politicalStability*0.1-5)
}

func updateMarketDemand(currentValue float64, tradeRelations float64) float64 {
	// Logic to update market demand based on trade relations
	return math.Max(0, currentValue+tradeRelations*0.1-5)
}

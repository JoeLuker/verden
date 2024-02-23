// File: simulation/simulation.go
package simulation

import "math/rand"

func RunSimulation(params SimulationParams) SimulationResult {
    // Introduce random event possibility
    if rand.Float64() < 0.1 { // 10% chance of a random event
        handleRandomEvent(params)
    }

    params.Economic.Update()
    params.Skills.Improve()

    economyHealth := calculateEconomyHealth(params.Economic, params.Skills)
    marketTrends := calculateMarketTrends(params.Economic, params.Skills)

    return SimulationResult{
        EconomyHealth: economyHealth,
        MarketTrends:  marketTrends,
    }
}

func handleRandomEvent(params SimulationParams) {
    // Placeholder for random event logic
    // Example: params.Economic.ResourceAvailability *= 0.9
}

func calculateEconomyHealth(e *EconomicFactors, s *CharacterSkills) float64 {
    // Updated logic for calculating economy health
    return (e.ResourceAvailability + s.CraftingAbility) - e.MarketDemand
}

func calculateMarketTrends(e *EconomicFactors, s *CharacterSkills) float64 {
    // Updated logic for calculating market trends
    return e.PoliticalStability * s.Negotiation / 100
}

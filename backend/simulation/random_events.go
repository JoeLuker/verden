// File: simulation/random_events.go
package simulation

import "math/rand"

// RandomEvent represents a type of random event that can occur in the simulation.
type RandomEvent struct {
    EventType string
    Impact    func(*SimulationParams)
}

// GenerateRandomEvent generates a random event based on the current state of the simulation.
func GenerateRandomEvent(params *SimulationParams) *RandomEvent {
    // List of possible random events
    events := []RandomEvent{
        {"MarketCrash", marketCrash},
        {"ResourceDiscovery", resourceDiscovery},
        {"PoliticalTurmoil", politicalTurmoil},
    }

    // Select a random event
    selectedEvent := events[rand.Intn(len(events))]
    return &selectedEvent
}

// marketCrash represents an economic downturn event.
func marketCrash(params *SimulationParams) {
    params.Economic.MarketDemand *= 0.8
}

// resourceDiscovery represents a sudden increase in resources.
func resourceDiscovery(params *SimulationParams) {
    params.Economic.ResourceAvailability *= 1.2
}

// politicalTurmoil represents a decrease in political stability.
func politicalTurmoil(params *SimulationParams) {
    params.Economic.PoliticalStability *= 0.9
}

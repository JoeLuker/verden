package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoeLuker/verden/simulation"
)

func HandleSimulation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		log.Println("HandleSimulation: method not allowed")
		return
	}

	var params simulation.SimulationParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("HandleSimulation: invalid request body: %v\n", err)
		return
	}

	// Validate the simulation parameters
	if err := params.Validate(); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		log.Printf("HandleSimulation: validation error: %v\n", err)
		return
	}

	// Log the received request
	log.Printf("HandleSimulation: Received request: %v", params)

	// Run the simulation
	result := simulation.RunSimulation(params)

	// Set the content type and encode the result
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("HandleSimulation: error encoding response: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

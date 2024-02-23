// File: handlers/simulation_result_handler.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoeLuker/verden/db"
)

// HandleRetrieveSimulation handles requests to retrieve the latest simulation results.
// HandleRetrieveSimulation handles requests to retrieve the latest simulation results.
func HandleRetrieveSimulation(w http.ResponseWriter, r *http.Request, simDataPersistence *db.SimulationDataPersistence) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	result, err := simDataPersistence.RetrieveLatestSimulationResult(ctx)
	if err != nil {
		log.Printf("Error retrieving simulation result: %v\n", err)
		http.Error(w, "Failed to retrieve simulation result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Failed to encode simulation result: %v\n", err)
		http.Error(w, "Failed to process simulation result", http.StatusInternalServerError)
	}
}

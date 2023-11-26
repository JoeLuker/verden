package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/JoeLuker/verden/db"
	"github.com/JoeLuker/verden/simulation"
)

// SimulationHandler contains the dependencies for the simulation handlers.
type SimulationHandler struct {
	DataPersistence *db.SimulationDataPersistence
}

// NewSimulationHandler creates a new instance of SimulationHandler.
func NewSimulationHandler(dataPersistence *db.SimulationDataPersistence) *SimulationHandler {
	return &SimulationHandler{
		DataPersistence: dataPersistence,
	}
}

// HandleSimulation runs the simulation and handles the HTTP request.
func (sh *SimulationHandler) HandleSimulation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var params simulation.SimulationParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := params.Validate(); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	result := simulation.RunSimulation(params)

	// If you want to use the DataPersistence to save the result, uncomment and handle the error
	err := sh.DataPersistence.SaveSimulationResult(r.Context(), &result)
	if err != nil {
		http.Error(w, "Failed to save simulation result: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

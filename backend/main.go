package main

import (
	"encoding/json"
    "log"
    "net/http"
    "github.com/JoeLuker/verden/simulation" // Import your simulation package
)

func main() {
    // Setup your routes here
    http.HandleFunc("/start-simulation", startSimulationHandler)
    // Add other handlers as needed

    // Start the server
    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

// Handler to start the economic simulation
func startSimulationHandler(w http.ResponseWriter, r *http.Request) {
    var params simulation.SimulationParams
    err := json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := simulation.RunSimulation(params)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse, err := json.Marshal(result)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

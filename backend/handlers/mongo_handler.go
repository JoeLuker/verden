package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoeLuker/verden/db"
	"github.com/JoeLuker/verden/models"
)

func MongoNodeInsertHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
	var node models.DiagramNode

	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.InsertDocument(r.Context(), "DiagramNodes", node)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Value set successfully with ID: %v", result.InsertedID)))
}

func GetDiagramHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
    // Logic to retrieve the diagram
    diagram, err := service.GetDiagram(r.Context())
    if err != nil {
        // handle error
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Marshal the diagram into a JSON string
    diagramJSON, err := json.Marshal(diagram)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(diagramJSON)
}
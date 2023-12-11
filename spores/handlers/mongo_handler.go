package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoeLuker/verden/db"
	"github.com/JoeLuker/verden/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler for inserting a node into the database
func MongoNodeInsertHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
	var node models.DiagramNode

	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a new ObjectID
	node.ID = primitive.NewObjectID()

	result, err := service.InsertDocument(r.Context(), "DiagramNodes", node)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Node created successfully with ID: %v", result.InsertedID)))
}

// Handler for retrieving the entire diagram
func GetDiagramHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
	diagram, err := service.GetDiagram(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	diagramJSON, err := json.Marshal(diagram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(diagramJSON)
}

// Additional handler functions for links and other CRUD operations can be added here

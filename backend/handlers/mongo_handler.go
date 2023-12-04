package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoeLuker/verden/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

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

func MongoDiagramGetHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
	var node models.DiagramStructure

	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// Assuming 'node' has a field 'ID' which is used to find the specific node
	filter := bson.M{"_id": node.ID}
	result, err := service.FindDocument(r.Context(), "Diagram", filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Document not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetDiagramIDHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
	// Logic to retrieve the diagram ID
	diagramID, err := service.GetDiagramID(r.Context())
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(diagramID))
}

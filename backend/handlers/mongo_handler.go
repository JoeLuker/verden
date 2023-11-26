package handlers

import (
	// "context"
	// "encoding/json"
	"net/http"

	"github.com/JoeLuker/verden/db"
)

func mongoInsertHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
	// var document YourActualDocumentType // Replace with your actual struct
	// if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// _, err := service.InsertDocument(context.Background(), "YourCollectionName", document)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Document inserted successfully"))
}

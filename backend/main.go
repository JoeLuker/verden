package main

import (
	"encoding/json"
	"log"
	"net/http"
    "context"
	"github.com/JoeLuker/verden/db"
	"github.com/JoeLuker/verden/middleware"
	"github.com/JoeLuker/verden/simulation"
)

func handleSimulation(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var params simulation.SimulationParams
    if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := simulation.RunSimulation(params)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}



// Handler for setting a value in Redis
func redisSetHandler(w http.ResponseWriter, r *http.Request, service *db.RedisService) {

	ctx := r.Context() // Use the request's context

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.Set(ctx, payload.Key, payload.Value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Value set successfully"))
}

// Handler for getting a value from Redis
func redisGetHandler(w http.ResponseWriter, r *http.Request, service *db.RedisService) {

	ctx := r.Context() // Use the request's context

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, err := service.Get(ctx, key)
	if err == db.RedisNil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

// Example MongoDB handler function
// Uncomment and modify this function according to your actual document structure
// func mongoInsertHandler(w http.ResponseWriter, r *http.Request, service *db.MongoDBService) {
//     var document YourActualDocumentType // Replace with your actual struct
//     if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }
//     _, err := service.InsertDocument(context.Background(), "YourCollectionName", document)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
//     w.Write([]byte("Document inserted successfully"))
// }

func main() {
    ctx := context.Background()
    
	// Redis setup
	rdb := db.ConnectRedis(ctx)
	redisService := db.NewRedisService(rdb)

	// Setup Redis routes
	http.Handle("/redis-set", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redisSetHandler(w, r, redisService)
	})))
	http.Handle("/redis-get", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redisGetHandler(w, r, redisService)
	})))

	// MongoDB setup
	// mongoService := db.NewMongoDBService()

	// Setup MongoDB routes (if you have any specific routes for MongoDB)
	// Example:
	// http.HandleFunc("/mongo-insert", func(w http.ResponseWriter, r *http.Request) {
	//     mongoInsertHandler(w, r, mongoService) // Using mongoService
	// })

	// Setup the start-simulation route
    http.HandleFunc("/simulate", handleSimulation)

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

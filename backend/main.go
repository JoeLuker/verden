package main

import (
	"context"
	"log"
	"net/http"

	"github.com/JoeLuker/verden/db"
	"github.com/JoeLuker/verden/handlers"
	"github.com/JoeLuker/verden/middleware"
)

func main() {
	ctx := context.Background()

	// Redis setup
	rdb := db.ConnectRedis(ctx)
	redisService := db.NewRedisService(rdb)

	// MongoDB setup
	mongoService := db.NewMongoDBService()

	http.Handle("/api/create-diagram-node", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.MongoNodeInsertHandler(w, r, mongoService)
	})))

	http.Handle("/api/get-diagram", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.MongoDiagramGetHandler(w, r, mongoService)
	})))

	http.Handle("/api/get-diagram-id", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetDiagramIDHandler(w, r, mongoService)
	})))

	// Simulation data persistence setup
	simDataPersistence := db.NewSimulationDataPersistence(mongoService, redisService)

	// Setup Redis routes
	http.Handle("/api/redis-set", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RedisSetHandler(w, r, redisService)
	})))
	http.Handle("/api/redis-get", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RedisGetHandler(w, r, redisService)
	})))

	http.Handle("/api/form-structure", middleware.EnableCORS(http.HandlerFunc(handlers.FormStructureHandler)))

	// Setup the simulation route
	simHandler := handlers.NewSimulationHandler(simDataPersistence)
	http.Handle("/api/simulate", middleware.EnableCORS(http.HandlerFunc(simHandler.HandleSimulation)))

	// Setup the route to retrieve simulation results
	http.Handle("/api/simulation-result", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleRetrieveSimulation(w, r, simDataPersistence)
	})))

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

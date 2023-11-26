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
    
    // Simulation data persistence setup
    simDataPersistence := db.NewSimulationDataPersistence(mongoService, redisService)

    // Setup Redis routes
    http.Handle("/api/redis-set", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handlers.RedisSetHandler(w, r, redisService)
    })))
    http.Handle("/api/redis-get", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handlers.RedisGetHandler(w, r, redisService)
    })))

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

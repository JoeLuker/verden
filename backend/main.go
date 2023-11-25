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

	// Setup Redis routes
	http.Handle("/api/redis-set", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RedisSetHandler(w, r, redisService)
	})))
	http.Handle("/api/redis-get", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RedisGetHandler(w, r, redisService)
	})))

	// Setup the start-simulation route
	http.Handle("/api/simulate", middleware.EnableCORS(http.HandlerFunc(handlers.HandleSimulation)))
	
	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

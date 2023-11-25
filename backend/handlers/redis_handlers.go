package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoeLuker/verden/db"
)

// RedisSetHandler sets a value in Redis.
func RedisSetHandler(w http.ResponseWriter, r *http.Request, service *db.RedisService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := service.Set(r.Context(), payload.Key, payload.Value); err != nil {
		log.Printf("Error setting value in Redis: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Value set successfully"))
}

// RedisGetHandler retrieves a value from Redis.
func RedisGetHandler(w http.ResponseWriter, r *http.Request, service *db.RedisService) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, err := service.Get(r.Context(), key)
	if err == db.RedisNil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error retrieving value from Redis: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

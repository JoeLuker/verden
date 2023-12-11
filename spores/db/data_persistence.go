// File: db/data_persistence.go
package db

import (
	"context"
	"encoding/json"
	"log"

	"github.com/JoeLuker/verden/simulation"
)

// SimulationDataPersistence handles the persistence of simulation data.
type SimulationDataPersistence struct {
	MongoDB *MongoDBService
	Redis   *RedisService
}

// NewSimulationDataPersistence creates a new instance for handling simulation data.
func NewSimulationDataPersistence(mongoDB *MongoDBService, redis *RedisService) *SimulationDataPersistence {
	return &SimulationDataPersistence{
		MongoDB: mongoDB,
		Redis:   redis,
	}
}

// SaveSimulationResult saves the simulation result to both MongoDB and Redis.
func (s *SimulationDataPersistence) SaveSimulationResult(ctx context.Context, result *simulation.SimulationResult) error {
	// Serialize result for storage
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// Save to MongoDB
	_, err = s.MongoDB.InsertDocument(ctx, "simulation_results", result)
	if err != nil {
		log.Printf("Failed to save simulation result to MongoDB: %v\n", err)
		return err
	}

	// Save to Redis
	err = s.Redis.Set(ctx, "latest_simulation_result", string(resultJSON))
	if err != nil {
		log.Printf("Failed to cache simulation result in Redis: %v\n", err)
		return err
	}

	return nil
}

// RetrieveLatestSimulationResult retrieves the latest simulation result from Redis,
// falling back to MongoDB if not available in Redis.
func (s *SimulationDataPersistence) RetrieveLatestSimulationResult(ctx context.Context) (*simulation.SimulationResult, error) {
	// Attempt to retrieve from Redis
	resultJSON, err := s.Redis.Get(ctx, "latest_simulation_result")
	if err != nil {
		log.Printf("Failed to retrieve simulation result from Redis: %v\n", err)
		// Fallback to MongoDB
		return s.retrieveFromMongoDB(ctx)
	}

	// Deserialize result
	var result simulation.SimulationResult
	if err := json.Unmarshal([]byte(resultJSON), &result); err != nil {
		log.Printf("Failed to deserialize simulation result: %v\n", err)
		// Fallback to MongoDB
		return s.retrieveFromMongoDB(ctx)
	}

	return &result, nil
}

// retrieveFromMongoDB retrieves the latest simulation result from MongoDB.
func (s *SimulationDataPersistence) retrieveFromMongoDB(ctx context.Context) (*simulation.SimulationResult, error) {
	// Implement logic to retrieve the latest result from MongoDB
	// Placeholder: return nil, nil
	return nil, nil
}

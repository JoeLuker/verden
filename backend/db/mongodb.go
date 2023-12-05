package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JoeLuker/verden/models"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBService struct {
	Client       *mongo.Client
	DatabaseName string
}

// NewMongoDBService initializes MongoDBService with environment variables.
func NewMongoDBService() *MongoDBService {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("The MONGO_URI environment variable must be set.")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB successfully")

	databaseName := extractDatabaseName(mongoURI)

	return &MongoDBService{
		Client:       client,
		DatabaseName: databaseName,
	}
}

// extractDatabaseName parses the database name from the MONGO_URI.
func extractDatabaseName(uri string) string {
	uriParts := strings.Split(uri, "/")
	return uriParts[len(uriParts)-1]
}

// PersistToMongoDB saves data from Redis to MongoDB.
func (service *MongoDBService) PersistToMongoDB(ctx context.Context, redisClient *redis.Client, userID string) error {
	val, err := redisClient.Get(ctx, userID).Result()
	if err != nil {
		return err
	}

	var userData models.UserData
	if err := json.Unmarshal([]byte(val), &userData); err != nil {
		return err
	}

	collection := service.Client.Database(service.DatabaseName).Collection("users")
	_, err = collection.InsertOne(ctx, userData)
	return err
}

// RetrieveFromMongoDB gets user data from MongoDB and saves it to Redis.
func (service *MongoDBService) RetrieveFromMongoDB(ctx context.Context, redisService *RedisService, userID string) error {
	collection := service.Client.Database(service.DatabaseName).Collection("users")
	var userData models.UserData
	if err := collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userData); err != nil {
		return err
	}

	return redisService.SaveToRedis(ctx, &userData)
}

// InsertDocument adds a new document to the specified collection in MongoDB.
func (s *MongoDBService) InsertDocument(ctx context.Context, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := s.Client.Database(s.DatabaseName).Collection(collectionName)
	return collection.InsertOne(ctx, document)
}

// FindDocument retrieves a single document from the specified collection in MongoDB.
func (s *MongoDBService) FindDocument(ctx context.Context, collectionName string, filter bson.M) (bson.M, error) {
	collection := s.Client.Database(s.DatabaseName).Collection(collectionName)

	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		// Handle 'no document found' error separately if needed
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found: %v", err)
		}
		return nil, err
	}

	return result, nil
}

func (s *MongoDBService) GetDiagramID(ctx context.Context) (string, error) {
	log.Println("GetDiagramID: Start")
	collection := s.Client.Database("diagramDB").Collection("diagramStructures")
	var result models.DiagramStructure
	log.Println("GetDiagramID: Finding diagram ID")
	err := collection.FindOne(ctx, bson.M{"Nodes": bson.M{"$exists": true}}).Decode(&result)
	if err != nil {
		log.Printf("GetDiagramID: Error finding diagram ID: %v", err)
		return "", err
	}
	log.Printf("GetDiagramID: Found diagram ID: %s", result.ID.Hex()) // Now this line should work
	return result.ID.Hex(), nil
}

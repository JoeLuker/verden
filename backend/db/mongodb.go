package db

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "github.com/go-redis/redis/v8"
    "github.com/JoeLuker/verden/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type MongoDBService struct {
    Client       *mongo.Client
    DatabaseName string
}

// Initialize MongoDBService with environment variables
func NewMongoDBService() *MongoDBService {
    mongoURI := os.Getenv("MONGO_URI")

    if mongoURI == "" {
        log.Fatal("The MONGO_URI environment variable must be set.")
    }

    // Connect to MongoDB using the connection string in the MONGO_URI environment variable
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
	
    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }
    fmt.Println("Connected to MongoDB successfully")
	
    // Assuming the database name is part of the MONGO_URI and it's the path segment following the last '/'
    uriParts := strings.Split(mongoURI, "/")
    databaseName := uriParts[len(uriParts)-1]

    return &MongoDBService{
        Client:       client,
        DatabaseName: databaseName,
    }
}
	
// PersistToMongoDB saves the data from Redis to MongoDB
func (service *MongoDBService) PersistToMongoDB(ctx context.Context, redisClient *redis.Client, userID string) error {
	val, err := redisClient.Get(ctx, userID).Result()
	if err != nil {
		return err
	}

	var userData models.UserData
	err = json.Unmarshal([]byte(val), &userData)
	if err != nil {
		return err
	}

	collection := service.Client.Database(service.DatabaseName).Collection("users")
	_, err = collection.InsertOne(ctx, userData)
	return err
}

// RetrieveFromMongoDB gets user data from MongoDB and saves it to Redis
func (service *MongoDBService) RetrieveFromMongoDB(ctx context.Context, redisService *RedisService, userID string) error {
	collection := service.Client.Database(service.DatabaseName).Collection("users")
	var userData models.UserData
	err := collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userData)
	if err != nil {
		return err
	}

	return redisService.SaveToRedis(ctx, &userData)
}

// InsertDocument adds a new document to the specified collection in MongoDB
func (s *MongoDBService) InsertDocument(ctx context.Context, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := s.Client.Database(s.DatabaseName).Collection(collectionName)
	return collection.InsertOne(ctx, document)
}

// FindDocument retrieves a single document from the specified collection in MongoDB
func (s *MongoDBService) FindDocument(ctx context.Context, collectionName string, filter interface{}) *mongo.SingleResult {
	collection := s.Client.Database(s.DatabaseName).Collection(collectionName)
	return collection.FindOne(ctx, filter)
}
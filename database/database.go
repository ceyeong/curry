package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database : Main DB instance
var Database *mongo.Database

// InitDatabase : Initialize db instance
func InitDatabase() error {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_HOST")))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	defer cancel()
	if err != nil {
		return err
	}

	//Ping Database connection
	err = dbClient.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	//Get instance of db & return nil error
	Database = dbClient.Database(os.Getenv("MONGO_DB"))
	return nil
}

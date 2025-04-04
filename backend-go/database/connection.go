package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var TodoCollection *mongo.Collection

func Connect() {
    // Replace with your own URI or use Atlas
    uri := "mongodb://localhost:27017"
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("❌ Failed to connect to MongoDB:", err)
    }

    Client = client
    TodoCollection = client.Database("tododb").Collection("todos")
    log.Println("✅ Connected to MongoDB successfully")
}

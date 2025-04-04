package services

import (
    "context"
    "github.com/zephyr942/backend-go/database"
    "github.com/zephyr942/backend-go/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

func GetAllTodos() ([]models.Todo, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := database.TodoCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }

    var todos []models.Todo
    if err = cursor.All(ctx, &todos); err != nil {
        return nil, err
    }

    return todos, nil
}

func CreateTodo(todo models.Todo) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _, err := database.TodoCollection.InsertOne(ctx, todo)
    return err
}

func UpdateTodo(id string, updated models.Todo) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := database.TodoCollection.UpdateOne(ctx,
        bson.M{"_id": objID},
        bson.M{"$set": updated},
    )
    return err
}

func DeleteTodo(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := database.TodoCollection.DeleteOne(ctx, bson.M{"_id": objID})
    return err
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo is the model for a task
type Todo struct {
    ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Description string             `json:"description" bson:"description"`
    IsComplete  bool               `json:"isComplete" bson:"isComplete"`
}

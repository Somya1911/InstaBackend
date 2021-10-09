package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Posts struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Caption   string             `json:"caption" bson:"caption"`
	ImageUrL  string             `json:"imageurl" bson:"imageurl"`
	Timestamp string             `json:"timestamp" bson:"timestamp"`
}

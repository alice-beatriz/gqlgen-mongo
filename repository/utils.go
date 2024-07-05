package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

func convertId(id string) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic("Invalid ID")
	}
	return objectId
}

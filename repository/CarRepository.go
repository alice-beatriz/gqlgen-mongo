package repository

import (
	"context"

	"github.com/alice-beatriz/gqlgen-mongo/db"
	"github.com/alice-beatriz/gqlgen-mongo/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func carsRepository(ctx context.Context) *mongo.Collection {
	return db.Template(ctx, "persons")
}

func GetCars(ctx context.Context) ([]*model.Car, error) {
	cursor, err := carsRepository(ctx).Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	var res []*model.Car
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}
	return res, nil
}

func GetCar(ctx context.Context, id string) (*model.Car, error) {
	var res *model.Car

	err := carsRepository(ctx).FindOne(ctx, bson.M{"_id": convertId(id)}).Decode(&res)
	if err != nil {
		panic(err)
	}
	return res, nil
}

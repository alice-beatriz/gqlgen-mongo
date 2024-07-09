package model

import (
	"context"

	"github.com/alice-beatriz/gqlgen-mongo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Car struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Model   string             `json:"model"`
	OwnerID primitive.ObjectID `json:"owner"`
}

func (c *Car) Owner(ctx context.Context) *Person {
	person, err := GetPerson(ctx, c.OwnerID)
	if err != nil {
		panic(err)
	}
	return person
}

func carsRepository(ctx context.Context) *mongo.Collection {
	return db.Template(ctx, "cars")
}

func GetCars(ctx context.Context) ([]*Car, error) {
	cursor, err := carsRepository(ctx).Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	var res []*Car
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}
	return res, nil
}

func GetCar(ctx context.Context, id primitive.ObjectID) (*Car, error) {
	var res *Car

	err := carsRepository(ctx).FindOne(ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil {
		panic(err)
	}
	return res, nil
}

func CreateCar(ctx context.Context, car CarInput) (*Car, error) {
	var model = Car{
		Model:   car.Model,
		OwnerID: car.Owner,
	}

	res, err := carsRepository(ctx).InsertOne(ctx, model)
	if err != nil {
		panic(err)
	}

	model.ID = res.InsertedID.(primitive.ObjectID)

	return &model, nil
}

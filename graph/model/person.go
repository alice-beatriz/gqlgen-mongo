package model

import (
	"context"

	"github.com/alice-beatriz/gqlgen-mongo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
}

func personRepository(ctx context.Context) *mongo.Collection {
	return db.Template(ctx, "persons")
}

func GetPersons(ctx context.Context) ([]*Person, error) {
	cursor, err := personRepository(ctx).Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	var res []*Person
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}
	return res, nil
}

func GetPerson(ctx context.Context, id primitive.ObjectID) (*Person, error) {
	var res *Person

	err := personRepository(ctx).FindOne(ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil {
		panic(err)
	}
	return res, nil
}

func CreatePerson(ctx context.Context, person PersonInput) (*Person, error) {
	var model = Person{
		ID:      primitive.NewObjectID(),
		Name:    person.Name,
		Address: person.Address,
	}

	res, err := personRepository(ctx).InsertOne(ctx, model)
	if err != nil {
		panic(err)
	}
	model.ID = res.InsertedID.(primitive.ObjectID)

	return &model, nil
}

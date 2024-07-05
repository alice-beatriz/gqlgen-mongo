package repository

import (
	"context"

	"github.com/alice-beatriz/gqlgen-mongo/db"
	"github.com/alice-beatriz/gqlgen-mongo/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func personRepository(ctx context.Context) *mongo.Collection {
	return db.Template(ctx, "persons")
}

func GetPersons(ctx context.Context) ([]*model.Person, error) {
	cursor, err := personRepository(ctx).Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	var res []*model.Person
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}
	return res, nil
}

func GetPerson(ctx context.Context, id string) (*model.Person, error) {
	var res *model.Person

	err := personRepository(ctx).FindOne(ctx, bson.M{"_id": convertId(id)}).Decode(&res)
	if err != nil {
		panic(err)
	}
	return res, nil
}

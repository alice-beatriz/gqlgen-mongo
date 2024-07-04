package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri string = "mongodb://localhost:27017/"
var client *mongo.Client

func init() {
	conn, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	client = conn
}

func getDB(ctx context.Context) *mongo.Database {
	user := ctx.Value("user")
	return client.Database(user.(string))
}

func Template(ctx context.Context, coll string) *mongo.Collection {
	return getDB(ctx).Collection(coll)
}

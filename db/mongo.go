package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri string = "mongodb://localhost:27017/"
var client *mongo.Client

func Init() {
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	conn, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongoUri).SetMonitor(cmdMonitor))
	if err != nil {
		panic(err)
	}
	client = conn
}

func getDB(ctx context.Context) *mongo.Database {
	if client == nil {
		panic("MongoDB client is not initialized")
	}
	user := ctx.Value("user")
	return client.Database(user.(string))
}

func Template(ctx context.Context, coll string) *mongo.Collection {
	return getDB(ctx).Collection(coll)
}

func Disconnect() {
	if client != nil {
		client.Disconnect(context.Background())
	}
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

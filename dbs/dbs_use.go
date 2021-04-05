package dbs

import (
	"github.com/TraiOi/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Use(db string) *mongo.Database {
	client_options := options.Client().ApplyURI(db_url)
	client, err := mongo.Connect(ctx, client_options)
	if err != nil {
		util.ErrorLogger.Print(err)
	} else {
		util.InfoLogger.Print("Connected to MongoDB: ", db_url)
	}

	return client.Database(db)
}

package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

/*
NewConnection : Function description of the function, input, output, and notes if any
*/
func NewConnection(uri string, dbname string, timeout time.Duration) (err error) {

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	db = client.Database(dbname)

	return nil
}

package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Session *mongo.Session
	Client  *mongo.Client
	Context context.Context
}

func NewConnectionDB(dbname string, dbuser string, password string) *Connection {
	ctx := context.Background()
	url := fmt.Sprintf("mongodb+srv://%s:%s@mytestcluster.6d53q.mongodb.net/%s?retryWrites=true&w=majority", dbuser, password, dbname)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Print(err)
	}
	var session mongo.Session
	if session, err = client.StartSession(); err != nil {
		session.EndSession(ctx)
	}
	return &Connection{
		Session: &session,
		Client:  client,
		Context: ctx,
	}
}

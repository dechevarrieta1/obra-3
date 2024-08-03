package mongoutilsv1

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	err    error
	doOnce sync.Once
)

type MongoConnection struct {
	Url      string
	Database string
}

func (mc MongoConnection) MakeBasicConnection() (*mongo.Client, error) {

	// log.Printf("@1. NORMAL\n") // * This logs has been asked to be commented due to $$$ issues

	doOnce.Do(func() {
		clientOpts := options.Client().ApplyURI(mc.Url).SetMaxPoolSize(1000)

		client, err = mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			log.Println("@[MakeBasicConnection - MongoConnection] Err > ", err)
		}
	})

	// log.Println("@[MakeBasicConnection - MongoConnection] > Success") // * This logs has been asked to be commented due to $$$ issues

	return client, err
}

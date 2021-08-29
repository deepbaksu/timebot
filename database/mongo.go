package database

import (
	"context"
	"fmt"
	"github.com/deepbaksu/timebot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"strings"
	"time"
)

func obfuscate(mongoDBURI string) string {
	split := strings.Split(mongoDBURI, "://")

	if len(split) < 2 {
		return mongoDBURI
	}

	protocol := split[0]
	rest := split[1]

	split = strings.Split(rest, "@")

	if len(split) < 2 {
		return mongoDBURI
	}

	usernameAndPassword := split[0]
	rest = split[1]

	split = strings.Split(usernameAndPassword, ":")
	username := split[0]

	return fmt.Sprintf("%s://%s:xxxxxx@%s", protocol, username, rest)
}

func fatalExitIfMongoError(err error, mongoDBURI string) {
	if err != nil {
		log.Panicf("Failed to connect MongoDB. Please check $MONGODB_URI(%v) is a correct MongoDB URI (err => %v)", mongoDBURI, err)
	}
}

func ProvideMongoClient(cfg *config.Config) *mongo.Client {
	log.Printf("connecting to MONGODB_URI = %s", obfuscate(cfg.MongoDBURI))

	ctx, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel1()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDBURI).SetRetryWrites(false).SetWriteConcern(writeconcern.New(writeconcern.WMajority())))
	fatalExitIfMongoError(err, cfg.MongoDBURI)

	ctx, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	err = mongoClient.Ping(ctx, readpref.Primary())
	fatalExitIfMongoError(err, cfg.MongoDBURI)

	return mongoClient
}

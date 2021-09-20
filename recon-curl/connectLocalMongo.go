package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func main(){
	ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	//client, err:= mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err:= mongo.Connect(ctx,options.Client().ApplyURI("mongodb://useradmin:idon'tknow123@localhost/cities?retryWrites=true&w=majority"))
	if err != nil {
		log.Println("Unable to Connect to Mongo", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx,readpref.Primary())
	if err != nil {
		log.Println(err)
	}
	database, err := client.ListDatabaseNames(ctx,bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(database)

}

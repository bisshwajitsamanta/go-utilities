package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"reflect"
	"time"
)

type Person struct {
	Name string
	Age int
	city string
}

func main(){

	// Declare Context type object for managing multiple API requests
	ctx,cancel:= context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()

	// Declare host and port options to pass to Connect to the MongoDB and return Client instance

	client, err:= mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://cloud_user:HB9NsGkYCSoRmi5X@mongo.z7lnc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	fmt.Println("Client Options Type: ", reflect.TypeOf(client),"\n")
	if err != nil {
		log.Println(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}
	databases, err := client.ListDatabaseNames(ctx,bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	// Access a MongoDB collection through a database
	collection:= client.Database("mydb").Collection("persons")
	bishwajit:= Person{
		Name: "Bishwajit",
		Age:  32,
		city: "Durgapur",
	}
	insertResult,err := collection.InsertOne(ctx,bishwajit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single Document", insertResult.InsertedID)

	// Finding a document
	//var result Person
	//
	//err = collection.FindOne(ctx,bson.D{}).Decode(&result)
	//if err != nil {
	//	fmt.Println("FindOne() Error", err)
	//} else{
	//	fmt.Println("FindOne() result:",result)
	//	fmt.Println("FindOne() Name:",result.Name)
	//	fmt.Println("FindOne() Age:",result.Age)
	//	fmt.Println("FindOne() City:",result.city)
	//}

	//filterCursor, err:= collection.Find(ctx,bson.M{"city": "Durgapur"})
	//if err != nil {
	//	fmt.Println("Finding all documents ERROR:", err)
	//	defer filterCursor.Close(ctx)
	//} else{
	//	for filterCursor.Next(ctx){
	//		var result bson.M
	//		err := filterCursor.Decode(&result)
	//		if err != nil {
	//			fmt.Println("Cursor.Next() error", err)
	//		} else{
	//			fmt.Println("\nresult type:", reflect.TypeOf(result))
	//			fmt.Println("result:", result)
	//		}
	//	}
	//}
}

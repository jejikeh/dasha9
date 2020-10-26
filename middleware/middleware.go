package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var username = "<username>"
var host1 = "" // of the form foo.mongodb.net

func main() {

	ctx := context.TODO()

	mongoURI := fmt.Sprintf("mongodb+srv://q11ve:m312131L@cluster0.uiv9i.azure.mongodb.net/<dbname>?retryWrites=true&w=majority")
	fmt.Println("connection string is:", mongoURI)

	// Set client options and connect
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	fmt.Println("Connected to MongoDB!")
}

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Email string
	Age   int
}

func main() {
	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check if the connection was successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get a handle to the "mydatabase" database and the "persons" collection
	collection := client.Database("mydatabase").Collection("persons")

	// Insert a document
	person := Person{Name: "John Doe", Email: "johndoe@example.com", Age: 30}
	insertResult, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal("Failed to insert document:", err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	// Update a document
	filter := bson.M{"name": "John Doe"}
	update := bson.M{"$set": bson.M{"email": "john@example.com"}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("Failed to update document:", err)
	}
	fmt.Printf("Matched %v document(s) and updated %v document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Find a document
	var result Person
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal("Failed to find document:", err)
	}
	fmt.Println("Found document:", result)

	// Delete a document
	/*	deleteResult, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal("Failed to delete document:", err)
		}
		fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	*/
	// Disconnect from MongoDB
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}
	fmt.Println("Disconnected from MongoDB!")
}

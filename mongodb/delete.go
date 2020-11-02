package main

import (
	"context" // Manage multiple requests
	"fmt"     // Println() function
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"     // Log errors to terminal
	"reflect" // Get an object type
	"time"
)

func main() {
	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://59.57.13.147:27019")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions))

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error calling mongo.Connect():", err)
	}

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// Declare an instance of the database and collection
	db := client.Database("airdroid_kid")
	col := db.Collection("notifications")

	// *mongo.DeleteResult object returned by API call
	fmt.Println("Collection type:", reflect.TypeOf(col), "n")

	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex("5f17d260c6e65c2b4e4a9579")
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {

		// Call the DeleteOne() method by passing BSON
		res, err := col.DeleteOne(ctx, bson.M{"_id": idPrimitive})
		fmt.Println("DeleteOne Result TYPE:", reflect.TypeOf(res))

		if err != nil {
			log.Fatal("DeleteOne() ERROR:", err)
		} else {

			// Check if the response is 'nil'
			if res.DeletedCount == 0 {
				fmt.Println("DeleteOne() document not found:", res)
			} else {
				// Print the results of the DeleteOne() method
				fmt.Println("DeleteOne Result:", res)

				// *mongo.DeleteResult object returned by API call
				fmt.Println("DeleteOne TYPE:", reflect.TypeOf(res))
			}
		}
	}

	// BSON filter for all documents with a value of 1
	f := bson.M{"device_id": "ba341cfe1781450c1eb1f164c492256e", "updated_at": bson.M{"$lte": "2020-07-17 12:18:00"}}
	fmt.Println("\nDeleting documents with filter:", f)

	// Call the DeleteMany() method to delete docs matching filter
	res, err := col.DeleteMany(ctx, f)

	// *mongo.DeleteResult object returned by API call
	fmt.Println("DeleteMany() result TYPE:", reflect.TypeOf(res))

	// Check for DeleteMany() API errors
	if err != nil {
		log.Fatal("DeleteMany() ERROR:", err)
	} else {

		// Check if the response is 'nil'
		if res.DeletedCount == 0 {
			fmt.Println("DeleteMany() documents not found:", res)
		} else {
			// Print the result of DeleteMany()
			fmt.Printf("DeleteMany() TOTAL: %d", res.DeletedCount)
			fmt.Println("\nDeleteMany() result:", res)
		}
	}
}
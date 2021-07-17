package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

// Customer entity
type Customer struct {
	CustomerId string `json:"customerId"`
	Enabled    bool   `json:"enabled"`
	Name       string `json:"name"`
}

func main() {
	collection := "customer"
	gcpProject := "my-demo-project"

	ctx := context.Background()

	// Output if the Firestore emulator is being used
	if value, ok := os.LookupEnv("FIRESTORE_EMULATOR_HOST"); ok {
		log.Printf("Using Firestore Emulator: '%s'", value)
	}

	// Create Firestore client
	client, err := firestore.NewClient(ctx, gcpProject)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close connection to Firestore after we are done
	defer client.Close()

	customerId := "my-id"
	customer := Customer{
		CustomerId: customerId,
		Enabled:    true,
		Name:       "Jason",
	}

	// Create a document
	client.Collection(collection).Doc(customerId).Set(ctx, customer)
	// Retrieve the just created document
	doc, _ := client.Collection(collection).Doc(customerId).Get(ctx)
	// Print the document to veriy the result.
	log.Println(doc.Data())

}

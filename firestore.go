package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func writeADocument(client *firestore.Client, ctx context.Context, userMap map[string]interface{}, collection string) {
	_, _, err := client.Collection(collection).Add(ctx, userMap)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "hciware-spaceman"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

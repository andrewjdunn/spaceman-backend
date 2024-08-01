package main

import (
	"context"
)

func writeSomeUsers() {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	writeADocument(client, ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815}, "users")

	writeADocument(client, ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	}, "users")
}

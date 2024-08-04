package main

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"cloud.google.com/go/firestore"
)

const AccountInfoCollectionName string = "accountInfo"
const AccountNameFieldName string = "name"
const AccountPublicKeyFieldName string = "publicKey"
const AccountPrivateKeyFieldName string = "privateKey"

func addAccountHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Handle a Account add request %s", r.URL.RawQuery)

	if r.URL.Path != "/account/add" || r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	var accountName = r.URL.Query().Get("name")

	if !CheckIfNameIsValid(accountName) {
		http.Error(w, "Not valid account name", 400)
		return
	}

	ctx := context.Background()
	client := createClient(ctx)
	if !checkIfAccountExists(client, ctx, accountName) {

		var publicKey = r.URL.Query().Get("key")
		var privateKey = makePrivatekey()

		defer client.Close()

		writeADocument(client, ctx, map[string]interface{}{
			AccountNameFieldName:       accountName,
			AccountPublicKeyFieldName:  publicKey,
			AccountPrivateKeyFieldName: privateKey}, AccountInfoCollectionName)
	} else {
		http.Error(w, "Account already exists", 400)
	}
}

func checkIfAccountExists(client *firestore.Client, ctx context.Context, name string) bool {
	q := client.Collection(AccountInfoCollectionName).Select(AccountNameFieldName).
		Where(AccountNameFieldName, "==", name)
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error checking for existence of %s %v", name, err)
		return false
	}
	count := 0

	for _, _ = range docs {
		count += 1
	}
	return count > 0
}

func CheckIfNameIsValid(accountName string) bool {
	re := regexp.MustCompile(`^[a-zA-Z_0-9]+$`)
	return re.Match([]byte(accountName))
}

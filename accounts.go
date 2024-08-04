package main

import (
	"context"
	"log"
	"net/http"
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
	var publicKey = r.URL.Query().Get("key")
	var privateKey = makePrivatekey()

	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	writeADocument(client, ctx, map[string]interface{}{
		AccountNameFieldName:       accountName,
		AccountPublicKeyFieldName:  publicKey,
		AccountPrivateKeyFieldName: privateKey}, AccountInfoCollectionName)
}

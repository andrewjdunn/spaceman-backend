package main

import (
	"fmt"
	"log"
	"net/http"
)

func addAccountHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Handle a Account add request %s", r.URL.RawQuery)

	if r.URL.Path != "/account/add" || r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "This is a", r.Method, " request at ", r.URL.Path, " name is ", r.URL.Query().Get("name"))

}

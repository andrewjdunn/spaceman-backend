package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "This is not a web page")
}

func addHandlers() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/account/add", addAccountHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	addHandlers()
}

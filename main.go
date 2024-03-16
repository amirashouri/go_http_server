package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/amirashouri/go_http_server/api"
)

func main() {
	http.HandleFunc("/", api.Greeting)
	http.HandleFunc("/person", api.PersonHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

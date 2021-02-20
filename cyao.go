package main

import (
	"log"
	"net/http"

	ihttp "github.com/edwin/cyoa/internal/http"
)

func main() {
	log.Println("Hello world")

	http.HandleFunc("/history", ihttp.History)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

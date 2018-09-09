package main

import (
	"log"
	"net/http"

	"github.com/chrisgoffinet/embed-files/static"
)

func main() {
	log.Println("listening on :8080")
	http.Handle("/", http.FileServer(static.HTTP))
	http.ListenAndServe(":8080", nil)
}

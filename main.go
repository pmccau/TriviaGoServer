package main

import (
	"github.com/pmccau/TriviaGoServer/router"
	"net/http"
	"log"
	"fmt"
)


func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

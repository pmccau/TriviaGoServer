package main

import (
	"fmt"
	"github.com/pmccau/TriviaGoServer/game"
	"github.com/pmccau/TriviaGoServer/router"
	"log"
	"net/http"
	"os"
)


func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080...")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == ":" {
		port = ":8080"
	}


	game.StartGameServer()

	fmt.Println("PORT:", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

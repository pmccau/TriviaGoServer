package router

import (
	"encoding/json"
	"fmt"
	"github.com/pmccau/TriviaGoServer/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	x := `{"test": "this is a test"}`
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := json.MarshalIndent(x, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, string(resp))
	})

	router.HandleFunc("/api/test", middleware.Test)
	router.HandleFunc("/api/getQuestions", middleware.GetQuestions)
	router.HandleFunc("/api/getCategories", middleware.GetCategories)
	return router
}

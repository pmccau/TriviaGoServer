package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"github.com/pmccau/TriviaGoServer/middleware"
	"io/ioutil"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//// Read in the help page for the landing
	md, err := ioutil.ReadFile("assets/landing.md")
	if err != nil {
		panic(err)
	}
	html := blackfriday.MarkdownCommon(md)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(html))
	})

	router.HandleFunc("/api/test", middleware.Test)
	router.HandleFunc("/api/questions", middleware.GetQuestions)
	router.HandleFunc("/api/categories", middleware.GetCategories)
	router.HandleFunc("/api/help", middleware.HelpRequest)
	return router
}

package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
)

// Test api call method
func Test(w http.ResponseWriter, r *http.Request) {
	message := "this is a test from golang"
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var res interface{}
	_ = json.NewDecoder(r.Body).Decode(&res)
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
	parseQuestionRequest(res)
	json.NewEncoder(w).Encode(message)
}

// GetCategories will route the response for trivia categories, returning a list of
// data.Category elements
func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Request the questions from the API
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get("https://opentdb.com/api_category.php")
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()

	// Parse API response to get the questions (in the 'results'). It's a bit roundabout to
	// marshal then unmarshal, but couldn't get it to work using other methods
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)
	results := parsedResponse["trivia_categories"]
	categories := parseCategories(results)

	// Send the response back to the calling server
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(reflect.TypeOf(categories))
	fmt.Println(categories)
	fmt.Fprint(w, indentJSON(categories))
}

// GetQuestions will route a response of trivia questions from the source DB to the requester
// in the form of a Question array
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	// Request the questions from the API
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get("https://opentdb.com/api.php?amount=10&category=9&difficulty=easy")
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()

	// Parse API response to get the questions (in the 'results'). It's a bit roundabout to
	// marshal then unmarshal, but couldn't get it to work using other methods
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)
	results := parsedResponse["results"]
	questions := parseQuestions(results)

	// Send the response back to the calling server
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(questions)
}
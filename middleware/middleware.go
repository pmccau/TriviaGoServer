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
	// Receive the inbound request for questions w/ appropriate fields
	// Parse the request, then hit the Open Trivia DB for the questions
	var res interface{}
	_ = json.NewDecoder(r.Body).Decode(&res)
	req := parseQuestionRequest(res)
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get(req.RequestURL)
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()

	// Parse API response from the DB to get the questions (in the 'results')
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)
	results := parsedResponse["results"]
	questions := parseQuestions(results)

	// Send the response back to the calling server
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(questions)
}


func CreateLobby(w http.ResponseWriter, r *http.Request) {

}


func GetLobby(w http.ResponseWriter, r *http.Request) {
	type inboundRequest struct {
		Name 		string
		Passcode 	string
	}

	var client = &http.Client{Timeout: 10 *time.Second}
	apiRes, err := client.Get("https://opentdb.com/api_category.php")
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()
	var req inboundRequest
	json.NewDecoder(apiRes.Body).Decode(&req)
	jsonStr, err := json.Marshal(req)
	fmt.Print(jsonStr)
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
	SetResponseHeaders(w)
	fmt.Println(reflect.TypeOf(categories))
	fmt.Println(categories)
	fmt.Fprint(w, indentJSON(categories))
}

// GetQuestions will route a response of trivia questions from the source DB to the requester
// in the form of a Question array. This is the endpoint that requests should come into. They
// should then be parsed to build a request URL, which is sent to the DB, then the response
// is sent back
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	// Receive the inbound request for questions w/ appropriate fields
	// Parse the request, then hit the Open Trivia DB for the questions
	var res interface{}
	_ = json.NewDecoder(r.Body).Decode(&res)
	req := parseQuestionRequest(res)
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get(req.RequestURL)
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()

	// Parse API response from the DB to get the questions (in the 'results')
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)
	results := parsedResponse["results"]
	questions := parseQuestions(results)

	// Send the response back to the calling server
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(questions)
}

// HelpRequest will return a JSON response containing helpful info on how to use the API.
// It's basically documentation in lieu of real documentation
func HelpRequest(w http.ResponseWriter, r *http.Request) {
	// Send the response back to the calling server
	SetResponseHeaders(w)
	fmt.Fprint(w, indentJSON(getHelpResponse()))
}

func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
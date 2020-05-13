package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Question struct {
	Text 		string
	Answer		string
	Category 	string
	Difficulty  string
}

func NewQuestion(text string, answer string, category string, difficulty string) *Question {
	q := new(Question)
	q.Text = text
	q.Answer = answer
	q.Category = category
	q.Difficulty = difficulty
	return q
}

type Category struct {
	Name 	string
	ID 		int
}

func NewCategory(id int, name string) *Category {
	c := new(Category)
	c.ID = id
	c.Name = name
	return c
}

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

type questionSet struct {
	Category		string
	CategoryID		int
	Difficulty		string
	AnsType			string
	NumQuestions	int
	Questions		[]*Question
}

func parseQuestionRequest(request interface{}) *questionSet {
	// Handle the []interface{} returned by response results
	switch result := request.(type) {
	case map[string]interface {}:
		qs := new(questionSet)
		qs.Category = interfaceToString(result["category"])
		qs.Difficulty = interfaceToString(result["category"])
		qs.AnsType = interfaceToString(result["category"])
		//qs.Questions = GetQuestions()
		qs.NumQuestions = len(qs.Questions)
		qs.CategoryID = interfaceToInt(result["category"])
		return qs
	}
	return nil
}

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
	//fmt.Println("Cats:", categories)

	// Send the response back to the calling server
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(reflect.TypeOf(categories))
	fmt.Println(categories)
	fmt.Fprint(w, formatResponse(categories))
	//json.NewEncoder(w).Encode(formatResponse(categories))
}

func formatResponse(response interface{}) string {
	resp, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(resp)
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

func parseCategories(results interface{}) []Category {
	var categories []Category

	// Handle the []interface{} returned by response results
	switch result := results.(type) {
	case []interface{}:
		for _, val := range result {
			// Iterate over each map returned
			switch val := val.(type) {
			case map[string]interface {}:
				name := interfaceToString(val["name"])
				id := interfaceToInt(val["id"])
				c := NewCategory(id, name)
				categories = append(categories, *c)
			}
		}
	}
	return categories
}

// parseQuestions is a helper that returns an array of pointers to questions
// and is meant to be used only in the GetQuestions function
func parseQuestions(results interface{}) []*Question {
	var questions []*Question

	// Handle the []interface{} returned by response results
	switch result := results.(type) {
	case []interface{}:
		for _, val := range result {

			// Iterate over each map returned
			switch val := val.(type) {
			case map[string]interface {}:
				category := interfaceToString(val["category"])
				answer := interfaceToString(val["correct_answer"])
				difficulty := interfaceToString(val["difficulty"])
				text := interfaceToString(val["question"])
				q := NewQuestion(text, answer, category, difficulty)
				questions = append(questions, q)
			default:
				fmt.Println("ERROR: Found type", reflect.TypeOf(result), "but expected map[string]interface {}")
			}
		}
	default:
		// Shouldn't happen
		fmt.Println("ERROR: Found type", reflect.TypeOf(result), "but expected []interface{}")
	}

	return questions
}

// interfaceToString is a quick helper to convert from an ambiguous string to a real one
func interfaceToString(toConvert interface{}) string {
	switch a := toConvert.(type) {
	case string:
		return a
	}
	return ""
}

func interfaceToInt(toConvert interface{}) int {
	switch a := toConvert.(type) {
	case int:
		return a
	case float64:
		return int(a)
	}
	return -1
}
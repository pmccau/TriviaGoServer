package middleware

import (
	"fmt"
	"github.com/pmccau/TriviaGoServer/data"
	"reflect"
)

// parseCategories will parse a response returned by the open trivia db into usable
// category objects
func parseCategories(results interface{}) []data.Category {
	var categories []data.Category
	// Handle the []interface{} returned by response results
	switch result := results.(type) {
	case []interface{}:
		for _, val := range result {
			// Iterate over each map returned
			switch val := val.(type) {
			case map[string]interface {}:
				c := new(data.Category)
				c.Name = interfaceToString(val["name"])
				c.ID = interfaceToInt(val["id"])
				categories = append(categories, *c)
			}
		}
	}
	return categories
}

// parseQuestions returns an array of Question objects by parsing the response
// from the open trivia db containing questions. This should take in the response
// send back from the actual query to trivia DB
func parseQuestions(results interface{}) []*data.Question {
	var questions []*data.Question
	// Handle the []interface{} returned by response results
	switch result := results.(type) {
	case []interface{}:
		for _, val := range result {
			// Iterate over each map returned
			switch val := val.(type) {
			case map[string]interface {}:
				q := new(data.Question)
				q.Category = interfaceToString(val["category"])
				q.Answer = interfaceToString(val["correct_answer"])
				q.Difficulty = interfaceToString(val["difficulty"])
				q.Text = interfaceToString(val["question"])
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

func buildQuestionRequestURL(qs *data.QuestionSet) string {
	root := "https://opentdb.com/api.php?"
	amount := fmt.Sprintf("amount=%d", qs.NumQuestions)
	category := fmt.Sprintf("&category=%v", qs.CategoryID)
	if qs.CategoryID < 0 {
		category = ""
	}
	difficulty := fmt.Sprintf("&difficulty=%v", qs.Difficulty)
	if qs.Difficulty == "" {
		difficulty = ""
	}
	questionType := fmt.Sprintf("&type=%v", qs.QuestionType)
	if qs.QuestionType == "" {
		questionType = ""
	}
	return fmt.Sprintf("%s%s%s%s%s", root, amount, category, difficulty, questionType)
}

// parseQuestionRequest will parse a question request coming from the react server
// such that it understand exactly which request to relay to the open trivia db
func parseQuestionRequest(request interface{}) *data.QuestionSet {
	// Handle the []interface{} returned by response results
	switch result := request.(type) {
	case map[string]interface {}:
		qs := new(data.QuestionSet)
		qs.Category = interfaceToString(result["category"])
		qs.Difficulty = interfaceToString(result["difficulty"])
		qs.QuestionType = interfaceToString(result["questionType"])
		qs.NumQuestions = interfaceToInt(result["numQuestions"])
		qs.RequestURL = buildQuestionRequestURL(qs)
		fmt.Print(qs)
		return qs
	}
	return nil
}
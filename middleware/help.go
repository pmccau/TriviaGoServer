package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getHelpResponse() *map[string]interface{} {
	data, err := ioutil.ReadFile("assets/help.json")
	if err != nil {
		panic(err)
	}


	//helpOutput := `{
	//	"root_url" : "https://triviagoserver.herokuapp.com/api",
    //    "endpoints" : [{
	//		"name" : "questions",
	//		"url" : "https://triviagoserver.herokuapp.com/api/getQuestions",
	//		"request_body_example": {
	//			"category" : 0,
	//			"question_type" : "any",
	//			"question_quantity" : 2,
	//			"difficulty": "easy"
	//		},
	//		"request_body_fields": {
	//			"category": "integer that should be derived from the 'ID' field of each Category",
	//			"question_type" : "string indicating question type: 'any', 'multiple', 'boolean'",
	//			"question_quantity" : "integer representing number of questions to return",
	//			"difficulty" : "string indicating question difficulty: 'any', 'easy', 'medium', 'hard'"
	//		},
	//		"request_body_fields": {
	//			"Text": "string containing body of the question",
	//			"Answer": "string containing the answer to the question",
	//			"Category": "string containing category of question",
	//			"Difficulty": "string containing difficulty of question"
	//		},
	//		"response_body_example": [
	//		  {
	//			"Text": "Virgin Trains, Virgin Atlantic and Virgin Racing, are all companies owned by which famous entrepreneur?",
	//			"Answer": "Richard Branson",
	//			"Category": "General Knowledge",
	//			"Difficulty": "easy"
	//		  },
	//		  {
	//			"Text": "How many furlongs are there in a mile?",
	//			"Answer": "Eight",
	//			"Category": "General Knowledge",
	//			"Difficulty": "easy"
	//		  }
	//		],
	//		"note": "The response_body_example is what would potentially be returned from the request_body_example"
	//	}],
	//	[{
	//		"name":
	//	}]
	//}`
	var output map[string]interface{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("Output:", output)
	return &output
}
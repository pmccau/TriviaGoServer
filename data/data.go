package data

type Category struct {
	Name 	string
	ID 		int
}

type Question struct {
	Text 		string
	Answer		string
	Category 	string
	Difficulty  string
}

type QuestionSet struct {
	Category     string
	CategoryID   int
	Difficulty   string
	QuestionType string
	NumQuestions int
	Questions    []*Question
	RequestURL	 string
}
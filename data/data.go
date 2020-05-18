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

func NewQuestion(Text string, Answer string, Category string, Difficulty string) *Question {
	q := new(Question)
	q.Text = Text
	q.Answer = Answer
	q.Category = Category
	q.Difficulty = Difficulty
	return q
}

// QuestionSet
type QuestionSet struct {
	Category     string
	CategoryID   int
	Difficulty   string
	QuestionType string
	NumQuestions int
	Questions    []*Question
	RequestURL	 string
}


func NewQuestionSet(CategoryID int, QuestionType string, NumQuestions int, Difficulty string) *QuestionSet {
	qs := new(QuestionSet)
	qs.CategoryID = CategoryID
	qs.QuestionType = QuestionType
	qs.NumQuestions = NumQuestions
	qs.Difficulty = Difficulty
	return qs
}

// Lobby has a number of rounds and an ID
type Lobby struct {
	Rounds 			int
	CurrentRound 	int
	LobbyID 		string
	Teams 			[]*Team
	Questions 		[]*QuestionSet
	Categories		[]*Category
	Passcode		string
}

func NewLobby(numRounds int) *Lobby {
	l := new(Lobby)
	l.Rounds = numRounds
	l.LobbyID = GenerateGuid()
	// Load the Questions w/ default properties
	for i := 0; i < l.Rounds; i++ {
		l.Questions[i] = NewQuestionSet(0, "Any", 10, "Any")
	}
	return l
}

// Team has a name and TeamID
type Team struct {
	Name 	string
	TeamID	string
}

func NewTeam(Name string) {
	t := new(Team)
	t.Name = Name
	t.TeamID = GenerateGuid()
}

//// Round has a round number and a QuestionSet
//type Round struct {
//	RoundNumber int
//	Questions 	[]*QuestionSet
//}
//
//func NewRound(RoundNumber int) *Round {
//	r := new(Round)
//	r.RoundNumber = RoundNumber
//}
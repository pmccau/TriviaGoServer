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
	NumQuestions	int
	CurrentRound 	int
	LobbyID 		string
	Teams 			[]*Team
	Questions 		[]*QuestionSet
	Categories		[]*Category
	Passcode		string
	Status 			LobbyStatus
}

type LobbyStatus string
type GameAction string

const(
	InLobby LobbyStatus = "InLobby"
	InRound LobbyStatus = "InRound"
	Scoring LobbyStatus = "Scoring"
	GameOver LobbyStatus = "GameOver"

	Start GameAction = "Start"
	Next GameAction = "Next"
	Pause GameAction = "Pause"
	End	GameAction = "End"
	Previous GameAction = "Previous"
)

func AddTeamToLobby(l *Lobby, t *Team) {
	l.Teams = append(l.Teams, t)
}

func NewLobby(numRounds int, numQuestions int, passcode string) *Lobby {
	l := new(Lobby)
	l.Rounds = numRounds
	if len(passcode) > 0 {
		l.Passcode = passcode
	}
	l.NumQuestions = numQuestions
	l.LobbyID = GenerateGuid()
	return l
}

// Team has a name and TeamID
type Team struct {
	Name 	string
	TeamID	string
	Answers [][]string
	Score 	int
	IsHost	bool
}

func NewTeam(Name string) *Team {
	t := new(Team)
	t.Name = Name
	t.TeamID = GenerateGuid()
	return t
}

type Response struct {
	TeamID 		string
	TeamAnswer	string
	Question 	*Question
	IsCorrect 	bool
}

func NewResponse(TeamID string, TeamAnswer string, Question *Question) *Response {
	r := new(Response)
	r.TeamID = TeamID
	r.TeamAnswer = TeamAnswer
	r.Question = Question
	r.IsCorrect = TeamAnswer == r.Question.Answer
	return r
}
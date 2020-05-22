package game

import (
	"fmt"
	"github.com/pmccau/TriviaGoServer/data"
)

var lobbies map[string]*data.Lobby

// CreateLobby initializes a new lobby
func CreateLobby() *data.Lobby {
	l := data.NewLobby(5, 10, "")
	fmt.Println("About to add:", l.LobbyID, "to map")
	lobbies[l.LobbyID] = l
	return l
}

// JoinLobby will attempt to add a new team to a lobby
func JoinLobby(LobbyID string, passcode string, Name string) *data.Lobby {
	if l, ok := lobbies[LobbyID]; ok {
		if l.Passcode == "" || l.Passcode == passcode {
			data.AddTeamToLobby(l, data.NewTeam(Name))
			return l
		}
	}
	return nil
}

// KillLobby will remove a lobby from the lobbies structure
func KillLobby(lobby data.Lobby) {
	if _, ok := lobbies[lobby.LobbyID]; ok {
		delete(lobbies, lobby.LobbyID)
	}
}

// PauseLobby will pause the lobby
func PauseLobby(lobby data.Lobby) {

}

// NextState will advance the lobby to the next state
func NextState(lobby data.Lobby) {
	switch lobby.Status {
	case data.Scoring:
		lobby.Status = data.InRound
		lobby.CurrentRound++
	case data.InLobby:
		lobby.Status = data.InRound
	case data.InRound:
		if lobby.CurrentRound == lobby.Rounds {
			lobby.Status = data.GameOver
		} else {
			lobby.Status = data.Scoring
			lobby.CurrentRound--
		}
	case data.GameOver:
		KillLobby(lobby)
	}
}

// PreviousState will move the lobby back to the most recent state
func PreviousState(lobby data.Lobby) {
	switch lobby.Status {
	case data.Scoring:
		lobby.Status = data.InRound
	case data.InRound:
		if lobby.CurrentRound == 1 {
			lobby.Status = data.InLobby
		} else {
			lobby.Status = data.Scoring
			lobby.CurrentRound--
		}
	case data.GameOver:
		lobby.Status = data.Scoring
	}
}

// HandleAction acts as router for actions
func HandleAction(Action data.GameAction, LobbyID string) {
	var f func(data.Lobby)
	if l, ok := lobbies[LobbyID]; ok {
		switch Action {
		case data.Start:
			f = NextState
		case data.Next:
			f = NextState
		case data.Previous:
			f = PreviousState
		case data.End:
			f = KillLobby
		case data.Pause:
			f = PauseLobby
		}
		f(*l)
	}
}

// StartGameServer begins a lobby and is a debugging test method
func StartGameServer() {
	lobbies = make(map[string]*data.Lobby)
	l := CreateLobby()
	fmt.Println("Lobby ID:", l.LobbyID)
	JoinLobby(l.LobbyID, "", "Pat")
	fmt.Println("Teams:", l.Teams[0])
}
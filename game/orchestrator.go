package game

import (
	"fmt"
	"github.com/pmccau/TriviaGoServer/data"
)

var lobbies map[string]*data.Lobby

type LobbyConfig struct {
	HostName string
	Passcode string
}

// CreateLobby initializes a new lobby
func CreateLobby(settings LobbyConfig) *data.Lobby {
	l := data.NewLobby(5, 10, settings.Passcode)
	fmt.Println("About to add:", l.LobbyID, "to map")
	lobbies[l.LobbyID] = l
	JoinLobby(l.LobbyID, settings.Passcode, settings.HostName)
	return l
}

// JoinLobby will attempt to add a new team to a lobby
func JoinLobby(lobbyID string, passcode string, name string) *data.Lobby {
	if l, ok := lobbies[lobbyID]; ok {
		if l.Passcode == "" || l.Passcode == passcode {
			data.AddTeamToLobby(l, data.NewTeam(name))
			return l
		}
	}
	return nil
}

// KillLobby will remove a lobby from the lobbies structure
func KillLobby(lobby data.Lobby, requester data.Team) {
	if _, ok := lobbies[lobby.LobbyID]; ok {
		delete(lobbies, lobby.LobbyID)
	}
}

// PauseLobby will pause the lobby
func PauseLobby(lobby data.Lobby, requester data.Team) {
	if requester.IsHost {
		lobby.Status = data.Paused
	}
}

// NextState will advance the lobby to the next state
// Valid transitions: 	Scoring (Rd n) => InRound (Rd n+1)
// 						Scoring (Rd max) => GameOver (Post-game)
//						InRound (Rd n) => Scoring (Rd n)
// 						InLobby (Pre-game) => InRound (Rd 1)
//						GameOver (Post-game) => Dismantle the lobby
func NextState(lobby data.Lobby, requester data.Team) {
	switch lobby.Status {
	case data.Scoring:
		if lobby.CurrentRound == lobby.Rounds {
			lobby.Status = data.GameOver
		} else {
			lobby.Status = data.InRound
			lobby.CurrentRound++
		}
	case data.InLobby:
		lobby.Status = data.InRound
	case data.InRound:
		lobby.Status = data.Scoring
		lobby.CurrentRound++
	case data.GameOver:
		KillLobby(lobby, requester)
	}
}

// PreviousState will move the lobby back to the most recent state
// Valid transitions:	Scoring (Rd n) => InRound (Rd n)
//						InRound (Rd n) => Scoring (Rd n-1)
// 						InRound (Rd 1) => InLobby (Pre-game)
//						GameOver (Post-game) => Scoring (Rd max)
func PreviousState(lobby data.Lobby, requester data.Team) {
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
func HandleAction(action data.GameAction, lobbyID string, requester data.Team) {
	var f func(data.Lobby, data.Team)
	if l, ok := lobbies[lobbyID]; ok {
		switch action {
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
		f(*l, requester)
	}
}

// StartGameServer begins a lobby and is a debugging test method
func StartGameServer() {
	lobbies = make(map[string]*data.Lobby)
	var settings LobbyConfig
	settings.HostName = "Player1-host"
	settings.Passcode = "abc"
	l := CreateLobby(settings)
	fmt.Println("Lobby created, ID:", l.LobbyID)
	JoinLobby(l.LobbyID, "abc", "Player2")
	fmt.Println("Teams:")
	for i, team := range l.Teams {
		fmt.Println(i, ". ", team.Name, "\tIsHost:", team.IsHost)
	}
}
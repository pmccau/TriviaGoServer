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
func KillLobby(LobbyID string) {
	if _, ok := lobbies[LobbyID]; ok {
		delete(lobbies, LobbyID)
	}
}

func UpdateScores(LobbyID string) *data.Lobby {
	if l, ok := lobbies[LobbyID]; ok {
		for _, t := range l.Teams {
			t.Score :=
		}
	}
	return nil
}

func StartRound(LobbyID string, RoundNum int) *data.Lobby {
	if l, ok := lobbies[LobbyID]; ok {
		switch l.Status {
		case data.GameOver:
			return l
		}
	}
	return nil 			// TEMPORARY
}

func ReturnToLobby(LobbyID string) *data.Lobby {

}

func TransitionLobbyState(l *data.Lobby, state data.LobbyStatus) *data.Lobby {
	if l.Status != state {
		switch state {
		case data.InRound:
			StartRound(l.LobbyID, l.CurrentRound + 1)
		case data.InLobby:
			print("a")
		case data.Scoring:
			print("a")
		case data.GameOver:
			KillLobby(l.LobbyID)
		}
	}
	return l
}

func StartGameServer() {
	lobbies = make(map[string]*data.Lobby)
	l := CreateLobby()
	fmt.Println("Lobby ID:", l.LobbyID)
	JoinLobby(l.LobbyID, "", "Pat")
	fmt.Println("Teams:", l.Teams[0])


}
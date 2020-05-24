package game

import (
	"testing"
)

func TestLobby(t *testing.T) {
	settings := new(LobbyConfig)
	settings.Passcode = "abc"
	settings.HostName = "P1"
	l := CreateLobby(*settings)
	JoinLobby(l.LobbyID, l.Passcode, "P2")
	if len(l.Teams) != 2 {
		t.Errorf("Number of teams = %d; want 2", len(l.Teams))
	}
}
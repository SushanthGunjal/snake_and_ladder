package model

import (
	"fmt"
)

type Player struct {
	Name     string `json:"name"`
	Position int64  `json:"position"`
	Won      bool   `json:"won,omitempty"`
}

func InitPlayer(name string, position int64) *Player {
	return &Player{
		Name:     name,
		Position: 0,
		Won:      false,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Name: %s, position: %d", p.Name, p.Position)
}

func AddPlayers(players []Player) []*Player {
	var result []*Player
	for i := 0; i < len(players); i++ {
		player := InitPlayer(players[i].Name, players[i].Position)
		result = append(result, player)
	}
	return result
}

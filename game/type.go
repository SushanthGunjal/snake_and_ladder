package game

import "github.com/SushanthGunjal/snake_and_ladder/model"

type Game struct {
	Board          *model.Board
	Dice           *model.Dice
	Players        []*model.Player
	TotalMoves     int64
	MovesStratergy []int64
	TotalPlayers   int64
}

package game

import "github.com/SushanthGunjal/snake_and_ladder/model"

func InitGame(numberOfSnakes, numberOfLadders, boardSize int64, players []model.Player, input model.Input) *Game {
	return &Game{
		Board:          model.InitBoard(boardSize, numberOfSnakes, numberOfLadders, input),
		Dice:           model.InitDice(6),
		Players:        model.AddPlayers(players),
		TotalMoves:     input.TotalMoves,
		TotalPlayers:   input.NumberOfPlayers,
		MovesStratergy: input.MovesStratergy,
	}
}

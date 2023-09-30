package game

import (
	"fmt"
)

func (game *Game) Play() {

	if len(game.Players) == 0 {
		fmt.Printf("no players joined.")
		return
	}
	totalNoOfMoves := 0
	for i := int64(0); i < game.TotalMoves; i++ {
		//fmt.Printf("move number %d and dice roll is %d \n", i, game.MovesStratergy[i])

		currentPlayer := game.Players[i%game.TotalPlayers]
		roll := game.MovesStratergy[i]

		oldPosition := currentPlayer.Position

		newPosition := currentPlayer.Position + roll
		if newPosition > game.Board.GetEndValue() {

			// he will have to play again
			newPosition = oldPosition

			fmt.Printf("Player %s, rolled a %d and moved from %d to %d \n", currentPlayer.Name, roll, oldPosition, newPosition)
			continue
		}

		currentPlayer.Position = game.Board.GetNewPosition(currentPlayer.Position + roll)
		fmt.Printf("Player %s, rolled a %d and moved from %d to %d \n", currentPlayer.Name, roll, oldPosition, newPosition)

		if currentPlayer.Position == game.Board.GetEndValue() {
			currentPlayer.Won = true
			fmt.Printf("%s Won!!!\n", currentPlayer)
		} else {
		}

		if len(game.Players) < 2 {
			fmt.Printf("Game finished!! total moves: %d\n", totalNoOfMoves)
			return
		}
	}
}

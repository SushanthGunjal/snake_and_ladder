package model

type Input struct {
	NumberOfPlayers int64    `json:"number_of_players"`
	BoardSize       int64    `json:"board_size"`
	NumberOfSnakes  int64    `json:"number_of_snakes"`
	NumberOfLadders int64    `json:"number_of_ladders"`
	SnakePos        []Snake  `json:"snake_pos"`
	LadderPos       []Ladder `json:"ladder_pos"`
	Players         []Player `json:"players"`
	TotalMoves      int64    `json:"total_moves"`
	MovesStratergy  []int64  `json:"moves_stratergy"`
}

package model

type Ladder struct {
	StartPoint int64 `json:"start_point"`
	EndPoint   int64 `json:"end_point"`
}

func InitLadder(startPoint, endPoint int64) *Ladder {
	return &Ladder{
		StartPoint: startPoint,
		EndPoint:   endPoint,
	}
}

func AddLadders(totalLadders int, input []Ladder) []Ladder {
	var ladders []Ladder

	for i := 0; i < totalLadders; i++ {
		ladder := Ladder{
			StartPoint: input[i].StartPoint,
			EndPoint:   input[i].EndPoint,
		}
		ladders = append(ladders, ladder)
	}

	return ladders
}

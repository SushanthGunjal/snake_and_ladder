package model

type Snake struct {
	StartPoint int64 `json:"start_point"`
	EndPoint   int64 `json:"end_point"`
}

func InitSnake(startPoint, endPoint int64) *Snake {
	return &Snake{
		StartPoint: startPoint,
		EndPoint:   endPoint,
	}
}

func AddSnakes(totalSnakes int, input []Snake) []Snake {
	var snakes []Snake

	for i := 0; i < totalSnakes; i++ {
		snake := Snake{
			StartPoint: input[i].StartPoint,
			EndPoint:   input[i].EndPoint,
		}
		snakes = append(snakes, snake)
	}

	return snakes
}

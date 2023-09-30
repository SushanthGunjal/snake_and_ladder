package model

import (
	"fmt"
	"log"
)

type Board struct {
	size   int64
	snakes []*Snake
	ladder []*Ladder
}

func InitBoard(size, noOfSnakes, noOfLadders int64, input Input) *Board {

	log.Println("Board Initializing started", input)

	b := &Board{
		size: size,
	}

	// map to store start and end positions of snakes
	snakeMap := make(map[string]bool)
	ladderMap := make(map[string]bool)

	for i := 0; i < int(noOfSnakes); i++ {
		start := input.SnakePos[i].StartPoint
		end := input.SnakePos[i].EndPoint
		if end >= start || start == size {
			continue
		}
		if _, ok := snakeMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
			b.snakes = append(b.snakes, InitSnake(start, end))
			snakeMap[fmt.Sprintf("%d:%d", start, end)] = true
		}
	}

	log.Println("Board Initialized with snake", snakeMap)

	for i := 0; i < int(noOfLadders); i++ {

		start := input.LadderPos[i].StartPoint
		end := input.LadderPos[i].EndPoint
		if start >= end || start == 1 {
			continue
		}
		if _, ok := ladderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
			b.ladder = append(b.ladder, InitLadder(start, end))
			ladderMap[fmt.Sprintf("%d:%d", start, end)] = true
		}

	}

	log.Println("Board Initialized with ladders", ladderMap)

	log.Println("Board Initialized", b)
	return b
}

func (b *Board) isLadder(pos int64) (bool, int64) {
	for _, val := range b.ladder {
		if val.StartPoint == pos {
			return true, val.EndPoint
		}
	}
	return false, -1
}

func (b *Board) isSnake(pos int64) (bool, int64) {
	for _, val := range b.snakes {
		if val.StartPoint == pos {
			return true, val.EndPoint
		}
	}
	return false, -1
}

func (b *Board) GetEndValue() int64 {
	return b.size
}

func (b *Board) GetNewPosition(pos int64) int64 {
	if ok, val := b.isLadder(pos); ok {
		fmt.Printf("is Ladder. start: %d, end: %d\n", pos, val)
		return val
	}

	if ok, val := b.isSnake(pos); ok {
		fmt.Printf("is snake. start: %d, end: %d\n", pos, val)
		return val
	}

	return pos
}

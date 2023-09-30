package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_GetEndValue(t *testing.T) {
	b := Board{}
	b.size = 100

	result := b.GetEndValue()
	assert.Equal(t, int64(100), result)
}

func TestBoard_GetNewPosition(t *testing.T) {

	b := Board{
		size:   100,
		snakes: getTestingSnakes(),
		ladder: getTestingLadders(),
	}

	result1 := b.GetNewPosition(20)
	assert.Equal(t, int64(10), result1)

	result2 := b.GetNewPosition(30)
	assert.Equal(t, int64(19), result2)

	result3 := b.GetNewPosition(21)
	assert.Equal(t, int64(31), result3)

	result4 := b.GetNewPosition(22)
	assert.Equal(t, int64(43), result4)
}

func TestBoard_IsLadder(t *testing.T) {
	b := Board{
		ladder: getTestingLadders(),
	}

	result1, endPoint1 := b.isLadder(21)
	assert.Equal(t, result1, true)
	assert.Equal(t, endPoint1, int64(31))

	result2, endPoint2 := b.isLadder(50)
	assert.Equal(t, result2, false)
	assert.Equal(t, endPoint2, int64(-1))

}

func TestBoard_IsSnake(t *testing.T) {
	b := Board{
		snakes: getTestingSnakes(),
	}

	result1, endPoint1 := b.isSnake(20)
	assert.Equal(t, result1, true)
	assert.Equal(t, endPoint1, int64(10))

	result2, endPoint2 := b.isSnake(50)
	assert.Equal(t, result2, false)
	assert.Equal(t, endPoint2, int64(-1))

}

func getTestingSnakes() []*Snake {
	snakes := []*Snake{
		{
			StartPoint: 20,
			EndPoint:   10,
		},
		{
			StartPoint: 30,
			EndPoint:   19,
		},
	}

	return snakes
}

func getTestingLadders() []*Ladder {
	ladders := []*Ladder{
		{
			StartPoint: 21,
			EndPoint:   31,
		},
		{
			StartPoint: 22,
			EndPoint:   43,
		},
	}

	return ladders
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/SushanthGunjal/snake_and_ladder/game"
	"github.com/SushanthGunjal/snake_and_ladder/model"
	"io/ioutil"
	"log"
)

func main() {

	var input model.Input

	jsonContent, err := ioutil.ReadFile("input.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(jsonContent, &input)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println("input file is", input)

	snakeAndLadderGame := game.InitGame(input.NumberOfSnakes, input.NumberOfLadders, input.BoardSize, input.Players, input)

	snakeAndLadderGame.Play()

}

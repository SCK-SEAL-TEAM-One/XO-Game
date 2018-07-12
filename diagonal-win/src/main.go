package main

import (
	"xo"
	"fmt"
	"strconv"
)


func main() {
	game := xo.NewXOGame()
	var inputFromConsole string
	for ; game.IsContinueGame(); game.Round++{
		fmt.Print("Round", game.Round, "\nInput:")
		fmt.Scan(&inputFromConsole)
		number,_ := strconv.Atoi(inputFromConsole)
		if game.Round%2 == 0 {
			game.InputPlayerO(number)
		}else{
			game.InputPlayerX(number)
		}		
	}
	fmt.Println("Result:", game.GetWinner())
}
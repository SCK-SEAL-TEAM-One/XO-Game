package main

import (
	"fmt"
	"xo"
)

func main() {
	xogame := xo.XOGame{}
	var playerName1, playerName2 string
	fmt.Scanf("%s %s", &playerName1, &playerName2)
	fmt.Println(xogame.SetPlayerSequency(playerName1, playerName2))

	var currentPlayerName string
	var positionNumber int
	for step := 0; step <= xo.GAMEONEROUND; step++ {
		fmt.Scanf("%s %d", &currentPlayerName, &positionNumber)
		chooseResult := xogame.ChoosePosition(currentPlayerName, positionNumber)
		fmt.Println(chooseResult)
		if xogame.GetPlayerWin(xogame.Step) {
			fmt.Println("WIN!")
			break
		}
	}
}

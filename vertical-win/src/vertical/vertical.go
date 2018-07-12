package vertical

import (
	"model"
)

func DuplicateMark(row int, column int, board []model.Cell) bool {
	for _, cellValue := range board {
		if cellValue.Row == row && cellValue.Column == column {
			return true
		}
	}
	return false
}

func Mark(player string, row, column int) model.Cell {
	playerCell := model.Cell{Row: 1, Column: 1, MarkedPlayer: "Player1"}

	return playerCell
}

func GetWinner(board []model.Cell) string {
	vertical1Player1 := 0
	vertical2Player1 := 0
	vertical3Player1 := 0
	vertical1Player2 := 0
	vertical2Player2 := 0
	vertical3Player2 := 0

	for _, cellValue := range board {
		if cellValue.MarkedPlayer == "Player1" {
			if cellValue.Column == 1 {
				vertical1Player1++
				if vertical1Player1 == 3 {
					return cellValue.MarkedPlayer
				}
			} else if cellValue.Column == 2 {
				vertical2Player1++
				if vertical2Player1 == 3 {
					return cellValue.MarkedPlayer
				}
			} else if cellValue.Column == 3 {
				vertical3Player1++
				if vertical3Player1 == 3 {
					return cellValue.MarkedPlayer
				}
			}
		} else if cellValue.MarkedPlayer == "Player2" {
			if cellValue.Column == 1 {
				vertical1Player2++
				if vertical1Player2 == 3 {
					return cellValue.MarkedPlayer
				}
			} else if cellValue.Column == 2 {
				vertical2Player2++
				if vertical2Player2 == 3 {
					return cellValue.MarkedPlayer
				}
			} else if cellValue.Column == 3 {
				vertical3Player2++
				if vertical3Player2 == 3 {
					return cellValue.MarkedPlayer
				}
			}
		}
	}
	return ""
}

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
	playerCell := model.Cell{Row: row, Column: column, MarkedPlayer: player}

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
			if checkCellValueColumn(cellValue, 1) {

				increaseVerticalPlayer(&vertical1Player1)

				if vertical1Player1 == 3 {
					return cellValue.MarkedPlayer
				}
			}
			if checkCellValueColumn(cellValue, 2) {

				increaseVerticalPlayer(&vertical2Player1)

				if cheakPlayerwinner(&vertical2Player1) {
					return cellValue.MarkedPlayer
				}
			}
			if checkCellValueColumn(cellValue, 3) {

				increaseVerticalPlayer(&vertical3Player1)

				if cheakPlayerwinner(&vertical3Player1) {
					return cellValue.MarkedPlayer
				}
			}
		}

		if cellValue.MarkedPlayer == "Player2" {
			if checkCellValueColumn(cellValue, 1) {

				increaseVerticalPlayer(&vertical1Player2)

				if cheakPlayerwinner(&vertical1Player2) {
					return cellValue.MarkedPlayer
				}
			}
			if checkCellValueColumn(cellValue, 2) {
				increaseVerticalPlayer(&vertical2Player2)
				if cheakPlayerwinner(&vertical2Player2) {
					return cellValue.MarkedPlayer
				}
			}
			if checkCellValueColumn(cellValue, 3) {

				increaseVerticalPlayer(&vertical3Player2)

				if cheakPlayerwinner(&vertical3Player2) {
					return cellValue.MarkedPlayer
				}
			}
		}
	}
	return ""
}

func checkCellValueColumn(cellValue model.Cell, column int) bool {
	return cellValue.Column == column
}

func increaseVerticalPlayer(verticalPlayer *int) {
	*verticalPlayer += 1
}

func cheakPlayerwinner(verticalPlayer *int) bool {
	return *verticalPlayer == 3
}

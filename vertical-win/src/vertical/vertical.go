package vertical

import (
	"XO-Game/vertical-win/src/model"
)

func DuplicateMark(row int, column int, board []model.Cell) bool {
	for _, cellValue := range board {
		if cellValue.Row == row && cellValue.Column == column {
			return true
		}
	}
	return false
}

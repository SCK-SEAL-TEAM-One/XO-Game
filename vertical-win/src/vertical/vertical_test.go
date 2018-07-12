package vertical

import (
	"XO-Game/vertical-win/src/model"
	"testing"
)

func Test_DuplicateMark_Input_Row_1_Column_1_Should_Be_False(t *testing.T) {
	row := 1
	column := 1
	board := []model.Cell{}
	board = append(board, model.Cell{Row: 1, Column: 2, MarkedPlayer: "Player1"})
	expected := false

	actual := DuplicateMark(row, column, board)

	if expected != actual {
		t.Errorf("Should be %t but got %t", expected, actual)
	}
}

func Test_DuplicateMark_Input_Row_1_Column_1_Should_Be_True(t *testing.T) {
	row := 1
	column := 1
	board := []model.Cell{}
	board = append(board, model.Cell{Row: 1, Column: 1, MarkedPlayer: "Player1"})
	expected := true

	actual := DuplicateMark(row, column, board)

	if expected != actual {
		t.Errorf("Should be %t but got %t", expected, actual)
	}
}

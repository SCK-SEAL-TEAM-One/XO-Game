package vertical

import (
	"model"
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

func Test_Mark_Input_Player1_And_Row_1_And_Column_1_Should_Be_StructCell(t *testing.T) {
	player := "Player1"
	row := 1
	column := 1
	expected := model.Cell{Row: 1, Column: 1, MarkedPlayer: "Player1"}

	actual := Mark(player, row, column)

	if expected != actual {
		t.Errorf("should be %v but it is %v ", expected, actual)
	}
}

func Test_GetWinner_Input_Struct_Board_Should_Be_Player1(t *testing.T) {
	board := []model.Cell{}
	board = append(board, model.Cell{Row: 1, Column: 1, MarkedPlayer: "Player1"})
	board = append(board, model.Cell{Row: 1, Column: 2, MarkedPlayer: "Player2"})
	board = append(board, model.Cell{Row: 2, Column: 1, MarkedPlayer: "Player1"})
	board = append(board, model.Cell{Row: 2, Column: 2, MarkedPlayer: "Player2"})
	board = append(board, model.Cell{Row: 2, Column: 1, MarkedPlayer: "Player1"})

	expected := "Player1"

	actual := GetWinner(board)

	if expected != actual {
		t.Errorf("should be %s but it is %s ", expected, actual)
	}
}

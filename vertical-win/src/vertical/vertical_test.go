package vertical

import (
    "testing"
    "model"
)

func Test_Mark_Input_Player1_And_Row_1_And_Column_1_Should_Be_StructCell(t *testing.T){
    player := "Player1"
    row := 1
    column := 1
    expected := model.Cell{Row:1,Column:1,MarkedPlayer:"Player1"}

    actual := Mark(player,row,column)

    if expected != actual {
		t.Errorf("should be %v but it is %v ", expected, actual)
	}
}
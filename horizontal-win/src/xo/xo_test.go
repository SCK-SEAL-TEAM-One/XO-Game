package xo

import (
	"fmt"
	"testing"
)

func Test_ChoosePosition_Input_Jay_1_Aun_5_Should_be_Array_15(t *testing.T) {
	xo := XOGame{}
	expected := []int{1, 5}
	xo.ChoosePosition("Jay", 1)

	actualJay := xo.ChoosePosition("Aun", 5)

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actualJay) {
		t.Errorf("Shoud Be %v but got %v", expected, actualJay)
	}
}

func Test_SetPlayerSequency_Input_Jay_Aun_Should_Be_Struct_Players(t *testing.T) {
	playerName1 := "JAY"
	playerName2 := "AUN"
	expected := []Player{
		{
			Name:     "JAY",
			Symbol:   "O",
			Sequence: "FIRST",
		},
		{
			Name:     "AUN",
			Symbol:   "X",
			Sequence: "SECOND",
		},
	}

	actual := SetPlayerSequency(playerName1, playerName2)

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

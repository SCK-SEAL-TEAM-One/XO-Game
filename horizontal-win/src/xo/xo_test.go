package xo

import (
	"fmt"
	"testing"
)

func Test_ChoosePosition_Input_Jay_1_Aun_5_Should_be_Array_15(t *testing.T) {
	xo := XOGame{}
	expected := []int{1, 5}

	xo.ChoosePosition("JAY", 1)
	actual := xo.ChoosePosition("AUN", 5)

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("Shoud Be %v but got %v", expected, actual)
	}
}

func Test_SetPlayerSequency_Input_Jay_Aun_Should_Be_Struct_Players(t *testing.T) {
	xo := XOGame{}
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

	actual := xo.SetPlayerSequency(playerName1, playerName2)

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func Test_GetPlayerWin_Input_array_1_5_3_4_2_Should_Be_True(t *testing.T) {
	xo := XOGame{}
	step := []int{1, 5, 3, 4, 2}
	expected := true

	actual := xo.GetPlayerWin(step)

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

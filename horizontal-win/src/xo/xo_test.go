package xo

import (
	"fmt"
	"testing"
)

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

func Test_GetPlayerWin_Input_array_1_5_3_4_2_Should_Be_True(t *testing.T) {
	step := []int{1, 5, 3, 4, 2}
	expected := true

	actual := GetPlayerWin(step)

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

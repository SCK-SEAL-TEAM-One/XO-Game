package xo

import (
	"testing"
)

func Test_InputPlayerO_Input_5_Should_Be_True(t *testing.T){
	input := 5
	expected := true
	xo := NewXOGame() 
	actual := xo.InputPlayerO(input) 
	if expected != actual {

		t.Errorf("Should be %v but got %v",expected, actual)
	}
}

func Test_InputPlayerX_Input_7_Should_Be_True(t *testing.T){
	input := 7
	expected := true
	xo := NewXOGame() 
	actual := xo.InputPlayerX(input) 
	if expected != actual {
		t.Errorf("Should be %v but got %v",expected, actual)
	}
}

func Test_CheckWinner_Should_Be_False(t *testing.T){
	expected := false
	xo := NewXOGame()
	actual := xo.CheckWinner()
	if expected != actual {
		t.Errorf("Should be %v but got %v",expected, actual)
	}
}
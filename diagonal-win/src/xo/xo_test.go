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

func Test_GetWinner_Should_Be_O_Win(t *testing.T){
	expected := "O WIN"
	xo := NewXOGame()
	xo.InputPlayerO(5)
	xo.InputPlayerX(7)
	xo.InputPlayerO(1)
	xo.InputPlayerX(8)
	xo.InputPlayerO(9)
	xo.CheckWinner()
	actual := xo.GetWinner()
	if expected != actual {
		t.Errorf("Should be %v but got %v",expected, actual)
	}
}

func Test_IsContinueGame_Should_Be_True(t *testing.T) {
	expected := true
	xo := NewXOGame()
	actual := xo.IsContinueGame()

	if expected != actual {
		t.Errorf("Should be %v but got %v",expected, actual)
	}
} 

func Test_IsContinueGame_Should_Be_False(t *testing.T){
	expected := false
	xo := NewXOGame()
	xo.InputPlayerO(5)
	xo.InputPlayerX(7)
	xo.InputPlayerO(1)
	xo.InputPlayerX(8)
	xo.InputPlayerO(9)
	xo.Round = 5
	xo.CheckWinner()
	actual := xo.IsContinueGame()
	if expected != actual {
		t.Errorf("Should be %v but got %v",expected, actual)
	}
}


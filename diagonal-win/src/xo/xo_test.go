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

		t.Errorf("Should be %v but got %v %v",expected, actual, xo.Table)
	}
}
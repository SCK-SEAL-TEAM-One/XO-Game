package xo

import (
	"testing"
)

func Test_InputPlayerO_Input_5_Should_Be_True(t *testing.T){
	input := 5
	expected := true 
	actual := InputPlayerO() 
	if expected != actual {
		t.Errorf("Should be %v but got %v ",expected, actual )
	}
}
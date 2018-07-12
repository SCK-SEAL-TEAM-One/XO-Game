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

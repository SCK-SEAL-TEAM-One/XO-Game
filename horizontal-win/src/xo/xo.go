package xo

type XOGame struct {
	Step []int
}

func (xo *XOGame) ChoosePosition(playerName string, position int) []int {
	xo.Step = append(xo.Step, position)
	return xo.Step
}

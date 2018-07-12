package xo

type XOGame struct {
	Step []int
}
type Player struct {
	Name     string
	Symbol   string
	Sequence string
}

func (xo *XOGame) ChoosePosition(playerName string, position int) []int {
	for _, valueStep := range xo.Step {
		if valueStep == position {
			return xo.Step
		}
	}
	xo.Step = append(xo.Step, position)
	return xo.Step

}
func SetPlayerSequency(playerName1, playerName2 string) []Player {
	return []Player{
		{
			Name:     playerName1,
			Symbol:   "O",
			Sequence: "FIRST",
		},
		{
			Name:     playerName2,
			Symbol:   "X",
			Sequence: "SECOND",
		},
	}
}

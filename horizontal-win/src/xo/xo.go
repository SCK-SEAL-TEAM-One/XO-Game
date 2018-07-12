package xo

import (
	"sort"
	"strconv"
	"strings"
)

type XOGame struct {
	Step []int
}

type Player struct {
	Name     string
	Symbol   string
	Sequence string
}

const GAMEONEROUND = 9

func (xo *XOGame) ChoosePosition(playerName string, position int) []int {
	for _, valueStep := range xo.Step {
		if valueStep == position {
			return xo.Step
		}
	}
	xo.Step = append(xo.Step, position)
	return xo.Step

}
func (xo *XOGame) SetPlayerSequency(playerName1, playerName2 string) []Player {
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

func GetPlayerWin(step []int) bool {
	var oddstep []int
	var evenstep []int
	for index, value := range step {
		if (index+1)%2 == 1 {
			oddstep = append(oddstep, value)
		}
		if (index+1)%2 == 0 {
			evenstep = append(evenstep, value)
		}
	}

	sort.Ints(oddstep)
	sort.Ints(evenstep)
	oddstring := []string{}
	for _, oddvalues := range oddstep {
		oddstring = append(oddstring, strconv.Itoa(oddvalues))
	}

	evenstring := []string{}
	for _, evenvalues := range evenstep {
		evenstring = append(evenstring, strconv.Itoa(evenvalues))
	}

	odd := strings.Join(oddstring, "")
	even := strings.Join(evenstring, "")
	if strings.Contains(odd, "123") || strings.Contains(even, "123") {
		return true
	}
	if strings.Contains(odd, "456") || strings.Contains(even, "456") {
		return true
	}
	if strings.Contains(odd, "789") || strings.Contains(even, "789") {
		return true
	}

	return false
}

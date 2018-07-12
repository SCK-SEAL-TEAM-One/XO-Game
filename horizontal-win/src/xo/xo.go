package xo

type Player struct {
	Name     string
	Symbol   string
	Sequence string
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

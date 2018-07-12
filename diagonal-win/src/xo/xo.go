package xo

type XOGAME struct{
	Table [3][3]string
	Round int
	MaxRound int
}
 
func NewXOGame() XOGAME{
	return XOGAME{MaxRound:9} 
} 

type Cell struct {
	Row int
	Col int
}
func (xo *XOGAME) InputPlayerO(number int) bool{
	numberCell := map[int] Cell{
		1: Cell{ Row: 0, Col: 0 },
		2: Cell{ Row: 0, Col: 1 },
		3: Cell{ Row: 0, Col: 2 },
		4: Cell{ Row: 1, Col: 0 },
		5: Cell{ Row: 1, Col: 1 },
		6: Cell{ Row: 1, Col: 2 },
		7: Cell{ Row: 2, Col: 0 },
		8: Cell{ Row: 2, Col: 1 },
		9: Cell{ Row: 2, Col: 2 },
	} 
	
	row := numberCell[number].Row
	col := numberCell[number].Col
	if xo.Table[row][col] == "" {
		xo.Table[row][col] = "O"
		return true
	}
	return false

}
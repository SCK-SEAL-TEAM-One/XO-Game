package xo

type XOGAME struct{
	Table [3][3]string
	Round int
	MaxRound int
}
 
func NewXOGame() XOGAME{
	return XOGAME{MaxRound:9} 
} 

func InputPlayerO() bool{
	return false

}
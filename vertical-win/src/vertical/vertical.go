package vertical

import (
    "model"
)

func Mark(player string,row,column int) model.Cell{
    playerCell :=  model.Cell{Row:1,Column:1,MarkedPlayer:"Player1"}
    return  playerCell
}
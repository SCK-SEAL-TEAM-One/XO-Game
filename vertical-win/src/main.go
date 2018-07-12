package main

import (
    "vertical"
    "model"
    "fmt"
)

func main()  {
    board  :=[]model.Cell{}
    playerWinner := ""

    for playerWinner == "" {
        fmt.Println("Player 1 Enter Row: ")
        var row int
        fmt.Scanln(&row)
        fmt.Println("Player 1 Enter Column: ")
        var column int
        fmt.Scanln(&column)

        if vertical.DuplicateMark(row,column,board) == false{
            playerSelectCell := vertical.Mark("Player1",row,column)
            board = append(board,playerSelectCell)
            playerWinner = vertical.GetWinner(board)
        }

        fmt.Println("Player 2 Enter Row: ")
        var rowPlayer2 int
        fmt.Scanln(&rowPlayer2)
        fmt.Println("Player 2 Enter Column: ")
        var columnPlayer2 int
        fmt.Scanln(&columnPlayer2)

        if vertical.DuplicateMark(rowPlayer2,columnPlayer2,board) == false{
            playerSelectCell := vertical.Mark("Player2",rowPlayer2,columnPlayer2)
            board = append(board,playerSelectCell)
            playerWinner = vertical.GetWinner(board)
        }

    }

    fmt.Println(playerWinner)
}
package model

type Board struct{
    Cell []Cell
}

type Cell struct{
    Row int
    Column int
    MarkedPlayer string
}
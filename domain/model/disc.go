package model

type Disc int

const (
	Empty Disc = iota
	Dark
	Light
)

const (
	E = Empty
	D = Dark
	L = Light
)

func NewInitialBoard() *[][]Disc {
	return &INITIAL_BOARD
}

var INITIAL_BOARD = [][]Disc{
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, D, L, E, E, E},
	{E, E, E, L, D, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
}

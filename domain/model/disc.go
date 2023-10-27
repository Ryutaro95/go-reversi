package model

type Disc int

const (
	Empty Disc = iota
	Dark
	Light
)

const E = Empty
const D = Dark
const L = Light

var INITIAL_BOARD = [][]Disc{
	[]Disc{E, E, E, E, E, E, E, E},
	[]Disc{E, E, E, E, E, E, E, E},
	[]Disc{E, E, E, E, E, E, E, E},
	[]Disc{E, E, E, D, L, E, E, E},
	[]Disc{E, E, E, L, D, E, E, E},
	[]Disc{E, E, E, E, E, E, E, E},
	[]Disc{E, E, E, E, E, E, E, E},
	[]Disc{E, E, E, E, E, E, E, E},
}
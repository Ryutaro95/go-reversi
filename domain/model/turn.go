package model

import (
	"time"
)

type Turn struct {
	GameId    int
	TurnCount int
	NextDisc  Disc
	Move      int
	Board     [][]Disc
	EndAt     time.Time
}

func NewFirstTurn(gameId int, endAt time.Time) Turn {
	return Turn{
		GameId:   gameId,
		NextDisc: Dark,
		Board:    INITIAL_BOARD,
		EndAt:    endAt,
	}
}

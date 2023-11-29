package model

import (
	"time"
)

type Turn struct {
	ID        int64
	GameID    int64
	TurnCount int
	NextDisc  Disc
	Move      int
	Board     [][]Disc
	EndAt     time.Time
}

func NewFirstTurn(gameID int64, endAt time.Time) *Turn {
	return &Turn{
		GameID:   gameID,
		NextDisc: Dark,
		Board:    *NewInitialBoard(),
		EndAt:    endAt,
	}
}

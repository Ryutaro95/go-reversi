package usecase

import (
	"database/sql"
	"log"
	"time"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type GetLatestTurnUsecase interface {
	GetLatestTurn(turnCount string) *GetLatestTurnOutput
}

type GetLatestTurn struct {
	DB         *sql.DB
	GameRepo   repository.GameRepo
	TurnRepo   repository.TurnRepo
	SquareRepo repository.SquareRepo
	MoveRepo   repository.MoveRepo
}

type GetLatestTurnOutput struct {
	Id         int64
	GameId     int64
	TurnCount  int
	Board      [][]int
	NextDisc   *model.Disc
	EndAt      time.Time
	WinnerDisc model.Disc
}

func NewGetLatestTurn(gameRepo repository.GameRepo, turnRepo repository.TurnRepo, squareRepo repository.SquareRepo, moveRepo repository.MoveRepo) GetLatestTurnUsecase {
	return &GetLatestTurn{
		GameRepo:   gameRepo,
		TurnRepo:   turnRepo,
		SquareRepo: squareRepo,
		MoveRepo:   moveRepo,
	}
}

func (gt *GetLatestTurn) GetLatestTurn(turnCount string) *GetLatestTurnOutput {
	game, err := gt.GameRepo.FindLatest()
	if err != nil {
		log.Fatalf("GetLatestTurn() fail: %v", err)
	}

	turn, err := gt.TurnRepo.FindForGameIdAndTurnCount(game.ID, turnCount)
	if err != nil {
		log.Fatalf("GetLatestTurn() fail: %v", err)
	}

	// TODO:  ゲームの勝敗を判定する
	winnerDisc := model.Empty

	squares, err := gt.SquareRepo.FetchSquaresByTurnID(turn.ID)
	board := convertSquaresToInt(squares)
	if err != nil {
		log.Fatalf("GetLatestTurn() fail: %v", err)
	}
	move, err := gt.MoveRepo.FindByTurnId(turn.ID)
	if move == nil {
		log.Println(move)
	}
	if err != nil {
		log.Fatalf("GetLatestTurn() fail: %v", err)
	}

	return &GetLatestTurnOutput{
		TurnCount:  turn.TurnCount,
		Board:      board,
		NextDisc:   &turn.NextDisc,
		WinnerDisc: winnerDisc,
	}
}

func convertSquaresToInt(squares []*model.Square) [][]int {
	initBoard := *model.NewInitialBoard()
	board := make([][]int, len(initBoard))
	for i := range board {
		board[i] = make([]int, len(board))
	}
	for _, square := range squares {
		board[square.Y][square.X] = int(square.Disc)
	}
	return board
}

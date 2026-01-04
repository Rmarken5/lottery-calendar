package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rmarken5/lottery-calendar/model"
)

var (
	ErrDateNotFound = errors.New("calendar day not found for given date")
)

type (
	Repository interface {
		GetDistinctGameTypes(ctx context.Context) ([]model.GameType, error)
		FindWinnersForGameOnDate(ctx context.Context, gameType model.GameType, winningNumbersString string, date time.Time) ([]model.Winner, error)
		BulkInsertDays(ctx context.Context, days []model.DayForInsert) error
		GetPlayers(ctx context.Context) ([]model.PlayerWithNumber, error)
	}

	Database struct {
		Reader
		Writer
	}
	Reader struct {
		db *sqlx.DB
	}
	Writer struct {
		db *sqlx.DB
	}
)

func New(readerConn *sqlx.DB, writerConn *sqlx.DB) *Database {
	return &Database{
		Reader: Reader{db: readerConn},
		Writer: Writer{db: writerConn},
	}
}

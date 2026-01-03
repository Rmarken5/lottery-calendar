package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rmarken5/lottery-calendar/model"
)

var (
	ErrDateNotFound = errors.New("calendar day not found for given date")
)

type (
	Repository interface {
		GetCalendarPrizeForDate(ctx context.Context, CalendarID uuid.UUID, date time.Time) (model.Day, error)
		// Create a method that queries based on gametype and drawing numbers to see if there's are matches and return the prizes.
		// Update winning numbers table to contain column that is representation of the numbers so that it can be easily compared.
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

func New(readerConn *sql.DB, writerConn *sql.DB) *Database {
	return &Database{
		Reader: Reader{db: sqlx.NewDb(readerConn, "postgres")},
		Writer: Writer{db: sqlx.NewDb(writerConn, "postgres")},
	}
}

const getCalendarPrizeForDateQuery = `SELECT FROM DAY d INNER JOIN CALENDAR c on d.calendar_id = d.id where d.id = $1 and d.date $2`

func (r *Reader) GetCalendarPrizeForDate(ctx context.Context, CalendarID uuid.UUID, date time.Time) (model.Day, error) {

}

package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/rmarken5/lottery-calendar/model"
)

const insertDayOfCalendarQuery = "insert into day_of_calendar (calendar_id, date, prize) values ($1, $2, $3);"

func (w *Writer) BulkInsertDays(ctx context.Context, days []model.DayForInsert) error {

	var errs []error
	for _, day := range days {
		_, err := w.db.ExecContext(ctx, insertDayOfCalendarQuery, day.CalendarID, day.Date.UTC().Format(time.RFC3339), day.Prize)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil

}

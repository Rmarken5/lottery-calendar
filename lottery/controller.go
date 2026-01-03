package lottery

import (
	"context"
	"time"

	"github.com/rmarken5/lottery-calendar/model"
)

type (
	Logic interface {
		GetWinningNumbers(ctx context.Context, gameType model.GameType, date time.Time) ([]string, error)
	}
	Controller struct {
	}
)

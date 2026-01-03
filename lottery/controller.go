package lottery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/rmarken5/lottery-calendar/database/postgres"
	"github.com/rmarken5/lottery-calendar/model"
)

const lotteryURL = "https://www.palottery.pa.gov/Custom/uploadedfiles/winning-numbers-history/PastWinningNumbers.ashx?g=%d&y=%d"

var (
	NumbersInGameType = map[model.GameType]uint8{
		model.GamePickThreeEvening: uint8(3),
		model.GamePickFourEvening:  uint8(4),
	}
)

type (
	Logic interface {
		GetWinningNumbers(ctx context.Context, gameType model.GameType, date time.Time) ([]uint8, error)
		BulkInsertCalendarDays(ctx context.Context, daysInJsonBytes []byte) error
	}
	Controller struct {
		repo       postgres.Repository
		httpClient http.Client
	}

	DrawingGame struct {
		DrawingGameID           int    `json:"drawingGameID"`
		DrawingGamePageEkID     int    `json:"drawingGamePageEkID"`
		DrawingNumberID         int    `json:"drawingNumberID"`
		DrawingNumberDate       string `json:"drawingNumberDate"`
		DrawingNumber1          int    `json:"drawingNumber1"`
		DrawingNumber2          int    `json:"drawingNumber2"`
		DrawingNumber3          int    `json:"drawingNumber3"`
		DrawingNumber4          int    `json:"drawingNumber4"`
		DrawingNumber5          *int   `json:"drawingNumber5"`
		DrawingNumber6          *int   `json:"drawingNumber6"`
		DrawingNumber7          *int   `json:"drawingNumber7"`
		DrawingNumber8          *int   `json:"drawingNumber8"`
		DrawingNumber9          *int   `json:"drawingNumber9"`
		DrawingNumber10         *int   `json:"drawingNumber10"`
		DrawingNumber11         *int   `json:"drawingNumber11"`
		DrawingNumberPayoutData string `json:"drawingNumberPayoutData"`
	}
)

func New(repo postgres.Repository) *Controller {
	return &Controller{
		repo:       repo,
		httpClient: http.Client{},
	}
}

func (c Controller) BulkInsertCalendarDays(ctx context.Context, daysInJsonBytes []byte) error {

	var days []model.DayForInsert
	err := json.NewDecoder(bytes.NewReader(daysInJsonBytes)).Decode(&days)
	if err != nil {
		slog.Error("error in decoding bytes into days", "error", err)
		return err
	}

	err = c.repo.BulkInsertDays(ctx, days)
	if err != nil {
		slog.Error("error writing days to database", "err", err)
	}

	return nil
}

func (c Controller) GetWinningNumbers(ctx context.Context, gameType model.GameType, date time.Time) (model.LotteryNumbers, error) {

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		slog.Error("error creating location", "err", err)
	}
	year := date.In(loc).Year()

	url := fmt.Sprintf(lotteryURL, int(gameType), year)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		slog.Error("error getting response from lottery", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		slog.Error("got non-200 level response from lottery", "status code", resp.StatusCode, "status", resp.Status)
		return nil, errors.New("got non-199 level response from lottery")
	}

	var numbers []DrawingGame
	err = json.NewDecoder(resp.Body).Decode(&numbers)
	if err != nil {
		slog.Error("unable to unmarshal response into struct", "error", err)
		return nil, err
	}
	if len(numbers) == 0 {
		slog.Info("no lottery for this day")
		return nil, nil
	}
	return numbers[len(numbers)-1:][0].NumbersFromGameType(gameType), nil

}

func (d DrawingGame) NumbersFromGameType(gameType model.GameType) model.LotteryNumbers {
	winningNumbers := make(model.LotteryNumbers, NumbersInGameType[gameType])
	switch gameType {
	case model.GamePickThreeEvening:
		winningNumbers[0] = uint8(d.DrawingNumber1)
		winningNumbers[1] = uint8(d.DrawingNumber2)
		winningNumbers[2] = uint8(d.DrawingNumber3)
	case model.GamePickFourEvening:
		winningNumbers[0] = uint8(d.DrawingNumber1)
		winningNumbers[1] = uint8(d.DrawingNumber2)
		winningNumbers[2] = uint8(d.DrawingNumber3)
		winningNumbers[3] = uint8(d.DrawingNumber4)
	}
	return winningNumbers
}

package model

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
)

//go:generate go tool stringer -type=GameType

type (
	// GameType indicates what lottery drawing is used for a Calendar
	GameType int

	// Calendar groups a years worth of Day and pairs it with a GameType so the user knows what
	// drawing to look for winning numbers
	Calendar struct {
		ID       uuid.UUID
		Name     string
		GameType GameType
	}

	// Day contains the prize for the Calendar on the specific Date
	Day struct {
		ID uuid.UUID
		Calendar
		Date  time.Time
		Prize string
	}

	// PlayerNumbers is used to hold the winning numbers of a Player for a Calendar.
	PlayerNumbers struct {
		CalendarID      uuid.UUID
		WinningNumbers  LotteryNumbers
		NumbersAsString string
	}

	// Player plays Games
	Player struct {
		ID            uuid.UUID
		Name          string
		PlayerNumbers []PlayerNumbers
	}
	Winner struct {
		PlayerName        string `db:"player_name"`
		PlayerPhoneNumber string `db:"phone_number"`
		PlayerEmail       string `db:"email"`
		CalendarName      string `db:"calendar_name"`
		Prize             string `db:"prize"`
	}

	DayForInsert struct {
		ID         uuid.UUID `json:"id,omitempty"`
		CalendarID uuid.UUID `json:"calendar_id,omitempty"`
		Date       time.Time `json:"date"`
		Prize      string    `json:"prize,omitempty"`
	}
)

const (
	GamePickFourEvening  GameType = 29 // PICK_FOUR_EVENING
	GamePickThreeEvening GameType = 28 // PICK_THREE_EVENING
)

func (gt *GameType) Scan(value driver.Value) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("value type unsupported for game type")
	}

	switch s {
	case GamePickThreeEvening.String():
		*gt = GamePickThreeEvening
		return nil
	case GamePickFourEvening.String():
		*gt = GamePickFourEvening
		return nil
	}

	return errors.New("unsupported game type from value")
}

func (gt GameType) Value() (driver.Value, error) {
	switch gt {
	case GamePickFourEvening:
		return GamePickFourEvening.String(), nil
	case GamePickThreeEvening:
		return GamePickThreeEvening.String(), nil
	}
	return "", errors.New("unsupported value from GameType")
}

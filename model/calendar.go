package model

import (
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
)

const (
	GamePickFourEvening  GameType = 29 // PICK_FOUR_EVENING
	GamePickThreeEvening GameType = 28 // PICK_THREE_EVENING
)

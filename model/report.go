package model

import "github.com/google/uuid"

type (
	DrawingReport struct {
		GameName     string
		Prize        string
		NumbersDrawn string
		Winner       *Player
	}

	PlayerWithNumber struct {
		PlayerID          uuid.UUID
		CalendarID        uuid.UUID
		PlayerName        string
		PlayerPhoneNumber string
		PlayerEmail       string
		Numbers           LotteryNumbers
	}
)

package postgres

import (
	"context"
	"time"

	"github.com/rmarken5/lottery-calendar/model"
)

const getDistinctGameTypesQuery = `select distinct game_type from calendar;`

func (r *Reader) GetDistinctGameTypes(ctx context.Context) ([]model.GameType, error) {
	var gameTypes = make([]model.GameType, 0)
	err := r.db.QueryRowContext(ctx, getDistinctGameTypesQuery).Scan(&gameTypes)
	if err != nil {
		return nil, err
	}
	return gameTypes, nil
}

const findWinnersForGameOnDateQuery = `select p.name as player_name,
p.phone_number, p.email,
c.name as calendar_name,
doc.prize
from player_numbers pn
inner join player p on p.id = pn.player_id
inner join calendar c on c.id = pn.calendar_id
inner join day_of_calendar doc on doc.calendar_id = c.id
where c.game_type = $1
and pn.winning_numbers_text = $2
and doc.date = $3;`

func (r *Reader) FindWinnersForGameOnDate(ctx context.Context, gameType model.GameType, winningNumbers string, date time.Time) ([]model.Winner, error) {
	winners := make([]model.Winner, 0)

	err := r.db.SelectContext(ctx, &winners, findWinnersForGameOnDateQuery, gameType.String(), winningNumbers, date.UTC().Format(time.DateOnly))
	if err != nil {
		return nil, err
	}

	return winners, nil
}

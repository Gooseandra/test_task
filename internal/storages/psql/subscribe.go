package psql

import (
	"context"
	"database/sql"
)

const insertTicker = `insert into "watching_coins"("name")values ($1) returning "id"`
const selectWatchingCoins = `select "name" from "watching_coins"`

type psqlSubscribes struct {
	db *sql.DB
}

func NewPsqlSubscribe(db *sql.DB) *psqlSubscribes { return &psqlSubscribes{db: db} }

func (p psqlSubscribes) Create(ctx context.Context, coinName string) (int, error) {
	var id int
	err := p.db.QueryRowContext(ctx, insertTicker, coinName).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (p psqlSubscribes) GetSymbols() ([]string, error) {
	rows, err := p.db.Query(selectWatchingCoins)
	if err != nil {
		return nil, err
	}
	var coinSymbols []string
	for rows.Next() {
		var temp string
		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		coinSymbols = append(coinSymbols, temp)
	}
	return coinSymbols, nil
}

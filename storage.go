package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const insertData = `insert into "coin_price"("name","price","time")values($1,$2,$3)`
const insertTicker = `insert into "watching_coins"("name")values ($1)`
const selectWatchingCoins = `select "name" from "watching_coins"`
const selectFetch = `select "name", "price", "time" from "coin_price"where "name" = $1 and "time" > $2 and "time" <$3`

type fetchData struct {
	name  string
	price string
	time  string
}

func AddPostgres(db *sql.DB, data PriceData) error {
	_, err := db.Exec(insertData, data.Symbol, data.Price, time.Now())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func AddTickerPostgres(db *sql.DB, coin string) error {
	_, err := db.Exec(insertTicker, coin)
	if err != nil {
		return err
	}
	return nil
}

func GetWatchingCoins(db *sql.DB) ([]string, error) {
	rows, err := db.Query(selectWatchingCoins)
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

func GetFetch(db *sql.DB, symbol string, timeFrom time.Time, timeTo time.Time) ([]fetchData, error) {
	rows, err := db.Query(selectFetch, symbol, timeFrom, timeTo)
	if err != nil {
		return nil, err
	}
	var data []fetchData
	for rows.Next() {
		var temp fetchData
		err = rows.Scan(&temp.name, &temp.price, &temp.time)
		if err != nil {
			return nil, err
		}
		data = append(data, temp)
	}
	fmt.Println(data)
	return data, nil
}

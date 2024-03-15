package psql

import (
	"database/sql"
	"fmt"
	"log"
	"test_task/internal/storages/model"
	"time"
)

const selectFetch = `select "name", "price", "time" from "coin_price"where "name" = $1 and "time" > $2 and "time" <$3`
const insertData = `insert into "coin_price"("name","price","time")values($1,$2,$3) returning "id"`

type psqlFetch struct {
	db *sql.DB
}

func NewPsqlFetchCoins(db *sql.DB) *psqlFetch { return &psqlFetch{db: db} }

func (f psqlFetch) FetchCoins(symbol string, timeFrom time.Time, timeTo time.Time) ([]model.FetchData, error) {
	rows, err := f.db.Query(selectFetch, symbol, timeFrom, timeTo)
	if err != nil {
		return nil, err
	}
	var data []model.FetchData
	for rows.Next() {
		var temp model.FetchData
		err = rows.Scan(&temp.Name, &temp.Price, &temp.Time)
		if err != nil {
			return nil, err
		}
		data = append(data, temp)
	}
	fmt.Println(data)
	return data, nil
}

func (f psqlFetch) NewFetchCoins(data model.PriceData) (int, error) {
	var id int
	err := f.db.QueryRow(insertData, data.Symbol, data.Price, time.Now()).Scan(&id)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	return id, nil
}

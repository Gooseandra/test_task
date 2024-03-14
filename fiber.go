package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type coinName struct {
	Name string `json:"name"`
}

type coinFetch struct {
	Symbol   string    `json:"ticker"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

type answer struct {
	Symbol      string
	StartPrice  float64
	FinishPrice float64
	Difference  float64
}

func fetch(app *fiber.App, db *sql.DB) {
	app.Get("/fetch", func(c *fiber.Ctx) error {
		var reqestData coinFetch
		err := json.Unmarshal(c.Body(), &reqestData)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}

		data, err := GetFetch(db, reqestData.Symbol, reqestData.DateFrom, reqestData.DateTo)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}

		priceStart, err := strconv.ParseFloat(data[0].price, 64)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}

		priceFinish, err := strconv.ParseFloat(data[len(data)-1].price, 64)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
		diff := (priceStart - priceFinish) / priceFinish * 100
		answerData := answer{Symbol: data[0].name, StartPrice: priceStart,
			FinishPrice: priceFinish, Difference: diff}
		jsonAnswer, err := json.Marshal(answerData)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
		return c.SendString(string(jsonAnswer))
	})
}

func add_ticker(app *fiber.App, db *sql.DB) {
	app.Post("/add_ticker", func(c *fiber.Ctx) error {
		var coin coinName
		err := json.Unmarshal(c.Body(), &coin)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
		err = checkCoinSymbol(coin.Name)
		if err != nil {
			return c.SendString(err.Error())
		}
		existingSymbols, err := GetWatchingCoins(db)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
		for _, item := range existingSymbols {
			if coin.Name == item {
				return c.SendString("This coin is already ticked")
			}
		}
		err = AddTickerPostgres(db, coin.Name)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
		go getFromBinance(coin.Name, db)
		return c.SendString("Coin " + coin.Name + " has been added")
	})
}

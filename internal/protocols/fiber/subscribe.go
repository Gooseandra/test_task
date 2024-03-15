package fiber

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test_task/internal/app"
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

type wrongSymbol struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func subscribe(c *fiber.Ctx, app app.App) (string, error) { //add_ticker

	var coin coinName
	err := json.Unmarshal(c.Body(), &coin)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	err = app.Binance.ValidCoin(coin.Name)
	if err != nil {
		return "", err
	}
	existingSymbols, err := app.Subscribe.GetSymbols()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	for _, item := range existingSymbols {
		if coin.Name == item {
			return "", errors.New("this coin is already ticked")
		}
	}
	id, err := app.Subscribe.Create(context.Background(), coin.Name)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	go app.Binance.GetFromBinance(coin.Name)
	return "Coin " + coin.Name + " has been added \nid: " + string(id), nil
}

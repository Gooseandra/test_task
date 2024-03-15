package fiber

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"test_task/internal/app"
)

type answer struct {
	Symbol      string
	StartPrice  float64
	FinishPrice float64
	Difference  float64
}

func fetchCoins(c *fiber.Ctx, app app.App) (string, error) {
	var reqestData coinFetch
	err := json.Unmarshal(c.Body(), &reqestData)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	data, err := app.Fetch.FetchCoins(reqestData.Symbol, reqestData.DateFrom, reqestData.DateTo)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if len(data) == 0 {
		return "", errors.New("no such data")
	}

	priceStart, err := strconv.ParseFloat(data[0].Price, 64)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	priceFinish, err := strconv.ParseFloat(data[len(data)-1].Price, 64)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	diff := (priceStart - priceFinish) / priceFinish * 100
	answerData := answer{Symbol: data[0].Name, StartPrice: priceStart,
		FinishPrice: priceFinish, Difference: diff}
	jsonAnswer, err := json.Marshal(answerData)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(jsonAnswer), nil
}

package servises

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"test_task/internal/storages/model"
	"time"
)

type Binance struct {
	url              string
	FetchStorage     model.FetchCoins
	SubscribeStorage model.Subscribes
}

func (b Binance) Init() error {
	symbols, err := b.SubscribeStorage.GetSymbols()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, item := range symbols {
		go b.GetFromBinance(item)
	}
	return nil
}

func NewBinance(url string, coins model.FetchCoins, subscribes model.Subscribes) *Binance {
	return &Binance{url: url, FetchStorage: coins, SubscribeStorage: subscribes}
}

type wrongSymbol struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (b Binance) ValidCoin(symbol string) error {
	resp, err := http.Get(b.url + symbol)
	if err != nil {
		fmt.Println("Checker: Ошибка при выполнении запроса:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Checker: Ошибка при чтении данных из ответа:", err)
		return err
	}
	var ws wrongSymbol
	err = json.Unmarshal(body, &ws)
	if err != nil {
		return err
	}
	fmt.Println(ws)
	if ws.Code != 0 {
		fmt.Println(ws.Code)
		return errors.New(ws.Msg)
	}
	return nil
}

func (b Binance) GetFromBinance(symbol string) {
	for {
		resp, err := http.Get(b.url + symbol)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка при чтении данных из ответа:", err)
			return
		}

		var price model.PriceData
		err = json.Unmarshal(body, &price)
		if err != nil {
			fmt.Println("Ошибка при парсинге JSON:", err)
			return
		}
		fmt.Printf("Символ: %s\nЦена: %f USDT\n", price.Symbol, price.Price)
		id, err := b.FetchStorage.NewFetch(price)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(id)
		time.Sleep(time.Minute)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type wrongSymbol struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func checkCoinSymbol(symbol string) error {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + symbol)
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

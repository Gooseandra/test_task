package model

import (
	"time"
)

type FetchCoins interface {
	FetchCoins(string, time.Time, time.Time) ([]FetchData, error)
	NewFetchCoins(PriceData) (int, error)
}

type FetchData struct {
	Name  string
	Price string
	Time  string
}

type PriceData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

package model

import (
	"time"
)

type FetchCoins interface {
	Fetch(string, time.Time, time.Time) ([]FetchData, error)
	NewFetch(PriceData) (int, error)
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

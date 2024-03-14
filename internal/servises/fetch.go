package servises

import (
	"test_task/internal/storages/model"
	"time"
)

type Fetch struct {
	storage model.FetchCoins
}

func NewFetch(storage model.FetchCoins) *Fetch { return &Fetch{storage: storage} }

func (f Fetch) Fetch(Symbol string, DateFrom time.Time, DateTo time.Time) ([]model.FetchData, error) {
	return f.storage.Fetch(Symbol, DateFrom, DateTo)
}

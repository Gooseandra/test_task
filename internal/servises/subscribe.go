package servises

import (
	"context"
	"test_task/internal/storages/model"
)

type Subscribe struct {
	storage model.Subscribes
}

func NewSubscribe(storage model.Subscribes) *Subscribe { return &Subscribe{storage: storage} }

func (s Subscribe) Create(ctx context.Context, coinName string) (int, error) {
	return s.storage.Create(ctx, coinName)
}

func (s Subscribe) GetSymbols() ([]string, error) { return s.storage.GetSymbols() }

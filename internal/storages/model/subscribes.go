package model

import "context"

type Subscribes interface {
	Create(ctx context.Context, coinName string) (int, error)
	GetSymbols() ([]string, error)
}

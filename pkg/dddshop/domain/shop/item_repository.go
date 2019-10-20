package shop

import "context"

type ItemRepository interface {
	Get(ctx context.Context, id string) (*Item, error)
	Put(ctx context.Context, src *Item) error
}

package shop

import (
	"context"
	"fmt"
)

type Service struct {
	itemRepository ItemRepository
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListItems(ctx context.Context) []*Item {
	return []*Item{{}}
}

type InputBuyItem struct {
	ItemID string
}

func (s *Service) BuyItem(ctx context.Context, input *InputBuyItem) (*Item, error) {
	//Transaction
	item, err := s.itemRepository.Get(ctx, input.ItemID)
	if err != nil {
		return nil, fmt.Errorf("payment service buy item err: %v", err)
	}

	if item.IsSold {
		return nil, &ItemIsAlreadySoldError{fmt.Errorf("payment service buy item err: item is alreay sold")}
	}

	item.IsSold = true
	if err := s.itemRepository.Put(ctx, item); err != nil {
		return nil, fmt.Errorf("payment service buy item err: %v", err)
	}

	return item, nil
}

package payment

import "context"

type AccountRepository interface {
	Get(ctx context.Context, accountID string) (*Account, error)
	GetByEmail(ctx context.Context, email string) (*Account, error)
	Put(ctxt context.Context, src *Account) error
}

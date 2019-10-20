package payment

import "context"

type ReceiptRepository interface {
	Put(ctx context.Context, src *Receipt) error
}

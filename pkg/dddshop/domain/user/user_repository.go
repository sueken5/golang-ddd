package user

import "context"

type UserRepository interface {
	Get(ctx context.Context, userID string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Put(ctx context.Context, src *User) error
}

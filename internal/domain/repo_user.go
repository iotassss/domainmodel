package domain

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByUUID(ctx context.Context, uuid UUID) (*User, error)
}

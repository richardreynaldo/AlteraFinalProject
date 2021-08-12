package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Address   string
	Email     string
	Password  string
	Role      string
	Token     string
}

type UseCase interface {
	CreateToken(ctx context.Context, email, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}

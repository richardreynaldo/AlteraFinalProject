package transaction_header

import (
	"context"
	"time"
)

type Domain struct {
	Id            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	UserId        int
	UserName      string
	TotalPrice    float64
	TotalQuantity int
	Status        string
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, transactionId int) (Domain, error)
	Store(ctx context.Context, coffeesDomain *Domain) (Domain, error)
	Update(ctx context.Context, coffeesDomain *Domain) (*Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, coffeesId int) (Domain, error)
	Store(ctx context.Context, coffeesDomain *Domain) (Domain, error)
	Update(ctx context.Context, coffeesDomain *Domain) (Domain, error)
	Find(ctx context.Context) ([]Domain, error)
}

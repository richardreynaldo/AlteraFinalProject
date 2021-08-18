package articles

import (
	"context"
	"time"
)

type Domain struct {
	Id          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	UserId      int
	UserName    string
	Description string
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, articlesId int) (Domain, error)
	GetByDescription(ctx context.Context, articlesDescription string) (Domain, error)
	Store(ctx context.Context, articlesDomain *Domain, userId int) (Domain, error)
	Update(ctx context.Context, articlesDomain *Domain) (*Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, articlesId int) (Domain, error)
	GetByDescription(ctx context.Context, articlesDescription string) (Domain, error)
	Store(ctx context.Context, articlesDomain *Domain) (Domain, error)
	Update(ctx context.Context, articlesDomain *Domain) (Domain, error)
	Find(ctx context.Context) ([]Domain, error)
}

package articles

import (
	articlesUsecase "finalProject/business/articles"
	"finalProject/drivers/databases/users"
	"time"
)

type Articles struct {
	Id          int
	Description string
	UserId      int
	User        users.User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain *articlesUsecase.Domain) *Articles {
	return &Articles{
		Id:          domain.Id,
		Description: domain.Description,
		UserId:      domain.UserId,
	}
}

func (rec *Articles) toDomain() articlesUsecase.Domain {
	return articlesUsecase.Domain{
		Id:          rec.Id,
		Description: rec.Description,
		UserId:      rec.UserId,
		UserName:    rec.User.Name,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

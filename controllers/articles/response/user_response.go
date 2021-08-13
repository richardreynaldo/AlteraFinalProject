package response

import (
	articles "finalProject/business/articles"
	"time"
)

type Article struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
	UserId      int       `json:"userId"`
	UserName    string    `json:"userName"`
	Description string    `json:"description"`
}

func FromDomain(domain articles.Domain) Article {
	return Article{
		Id:          domain.Id,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		UserId:      domain.UserId,
		UserName:    domain.UserName,
		Description: domain.Description,
	}
}

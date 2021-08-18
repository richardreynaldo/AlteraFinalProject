package response

import (
	reviews "finalProject/business/reviews"
	"time"
)

type Review struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	ArticleId int       `json:"article_id"`
	Review    string    `json:"review"`
	Rating    int       `json:"rating"`
}

func FromDomain(domain reviews.Domain) Review {
	return Review{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UserId:    domain.UserId,
		UserName:  domain.UserName,
		Rating:    domain.Rating,
		Review:    domain.Review,
		ArticleId: domain.ArticleId,
	}
}

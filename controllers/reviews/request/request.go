package request

import (
	"finalProject/business/reviews"
)

type CreateReview struct {
	Rating    int    `json:"rating"`
	Review    string `json:"review"`
	ArticleId int    `json:"article_id"`
}

func (req *CreateReview) ToDomain() *reviews.Domain {
	return &reviews.Domain{
		Rating:    req.Rating,
		Review:    req.Review,
		ArticleId: req.ArticleId,
	}
}

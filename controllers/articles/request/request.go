package request

import (
	"finalProject/business/articles"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateArticle struct {
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

func (req *CreateArticle) ToDomain() *articles.Domain {
	return &articles.Domain{
		Description: req.Description,
		UserId:      req.UserId,
	}
}

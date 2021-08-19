package reviews

import (
	reviewUsecase "finalProject/business/reviews"
	"finalProject/drivers/databases/articles"
	"finalProject/drivers/databases/users"
	"time"
)

type Reviews struct {
	Id        int
	Rating    int
	Review    string
	UserId    int
	User      users.User
	ArticleId int
	Article   articles.Articles
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt time.Time
}

func fromDomain(domain *reviewUsecase.Domain) *Reviews {
	return &Reviews{
		Id:        domain.Id,
		Rating:    domain.Rating,
		UserId:    domain.UserId,
		ArticleId: domain.ArticleId,
		Review:    domain.Review,
	}
}

func (rec *Reviews) toDomain() reviewUsecase.Domain {
	return reviewUsecase.Domain{
		Id:        rec.Id,
		UserId:    rec.UserId,
		UserName:  rec.User.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		Review:    rec.Review,
		Rating:    rec.Rating,
		ArticleId: rec.ArticleId,
	}
}

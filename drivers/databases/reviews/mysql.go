package reviews

import (
	"context"
	"finalProject/business/reviews"

	"gorm.io/gorm"
)

type mysqlReviewRepository struct {
	Conn *gorm.DB
}

func NewMySQLReviewRepository(conn *gorm.DB) reviews.Repository {
	return &mysqlReviewRepository{
		Conn: conn,
	}
}

func (ar *mysqlReviewRepository) Fetch(ctx context.Context, page, perpage int) ([]reviews.Domain, int, error) {
	rec := []Reviews{}

	offset := (page - 1) * perpage
	err := ar.Conn.Preload("User").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []reviews.Domain{}, 0, err
	}

	var totalData int64
	err = ar.Conn.Count(&totalData).Error
	if err != nil {
		return []reviews.Domain{}, 0, err
	}

	var domainNews []reviews.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (ar *mysqlReviewRepository) GetByID(ctx context.Context, newsId int) (reviews.Domain, error) {
	rec := Reviews{}
	err := ar.Conn.Where("id = ?", newsId).First(&rec).Error
	if err != nil {
		return reviews.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlReviewRepository) GetByDescription(ctx context.Context, newsTitle string) (reviews.Domain, error) {
	rec := Reviews{}
	err := ar.Conn.Where("description = ?", newsTitle).First(&rec).Error
	if err != nil {
		return reviews.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlReviewRepository) Store(ctx context.Context, newsDomain *reviews.Domain) (reviews.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Create(&rec)
	if result.Error != nil {
		return reviews.Domain{}, result.Error
	}
	// ini buat join
	err := ar.Conn.Preload("User").First(&rec, rec.Id).Error
	if err != nil {
		return reviews.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlReviewRepository) Update(ctx context.Context, newsDomain *reviews.Domain) (reviews.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Save(&rec)
	if result.Error != nil {
		return reviews.Domain{}, result.Error
	}

	err := ar.Conn.Preload("User").First(&rec, rec.Id).Error
	if err != nil {
		return reviews.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlReviewRepository) Find(ctx context.Context) ([]reviews.Domain, error) {
	rec := []Reviews{}

	query := ar.Conn.Preload("User")

	err := query.Find(&rec).Error
	if err != nil {
		return []reviews.Domain{}, err
	}

	articleDomain := []reviews.Domain{}
	for _, value := range rec {
		articleDomain = append(articleDomain, value.toDomain())
	}

	return articleDomain, nil
}

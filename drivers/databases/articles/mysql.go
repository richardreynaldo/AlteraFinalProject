package articles

import (
	"context"
	"finalProject/business/articles"

	"gorm.io/gorm"
)

type mysqlArticlesRepository struct {
	Conn *gorm.DB
}

func NewMySQLArticleRepository(conn *gorm.DB) articles.Repository {
	return &mysqlArticlesRepository{
		Conn: conn,
	}
}

func (ar *mysqlArticlesRepository) Fetch(ctx context.Context, page, perpage int) ([]articles.Domain, int, error) {
	rec := []Articles{}

	offset := (page - 1) * perpage
	err := ar.Conn.Preload("articles").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []articles.Domain{}, 0, err
	}

	var totalData int64
	err = ar.Conn.Count(&totalData).Error
	if err != nil {
		return []articles.Domain{}, 0, err
	}

	var domainNews []articles.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (ar *mysqlArticlesRepository) GetByID(ctx context.Context, newsId int) (articles.Domain, error) {
	rec := Articles{}
	err := ar.Conn.Where("id = ?", newsId).First(&rec).Error
	if err != nil {
		return articles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlArticlesRepository) GetByDescription(ctx context.Context, newsTitle string) (articles.Domain, error) {
	rec := Articles{}
	err := ar.Conn.Where("title = ?", newsTitle).First(&rec).Error
	if err != nil {
		return articles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlArticlesRepository) Store(ctx context.Context, newsDomain *articles.Domain) (articles.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Create(&rec)
	if result.Error != nil {
		return articles.Domain{}, result.Error
	}

	err := ar.Conn.Preload("Article").First(&rec, rec.Id).Error
	if err != nil {
		return articles.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlArticlesRepository) Update(ctx context.Context, newsDomain *articles.Domain) (articles.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Save(&rec)
	if result.Error != nil {
		return articles.Domain{}, result.Error
	}

	err := ar.Conn.Preload("Article").First(&rec, rec.Id).Error
	if err != nil {
		return articles.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlArticlesRepository) Find(ctx context.Context) ([]articles.Domain, error) {
	rec := []Articles{}

	query := ar.Conn.Preload("Article")

	err := query.Find(&rec).Error
	if err != nil {
		return []articles.Domain{}, err
	}

	articleDomain := []articles.Domain{}
	for _, value := range rec {
		articleDomain = append(articleDomain, value.toDomain())
	}

	return articleDomain, nil
}

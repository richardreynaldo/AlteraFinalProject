package transaction_detail

import (
	"context"
	"finalProject/business/transaction_detail"

	"gorm.io/gorm"
)

type mysqlTransactionDetailRepository struct {
	Conn *gorm.DB
}

func NewMySQLTransactionDetailRepository(conn *gorm.DB) transaction_detail.Repository {
	return &mysqlTransactionDetailRepository{
		Conn: conn,
	}
}

func (ar *mysqlTransactionDetailRepository) Fetch(ctx context.Context, page, perpage int) ([]transaction_detail.Domain, int, error) {
	rec := []TransactionDetail{}

	offset := (page - 1) * perpage
	err := ar.Conn.Preload("User").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []transaction_detail.Domain{}, 0, err
	}

	var totalData int64
	err = ar.Conn.Count(&totalData).Error
	if err != nil {
		return []transaction_detail.Domain{}, 0, err
	}

	var domainNews []transaction_detail.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (ar *mysqlTransactionDetailRepository) GetByID(ctx context.Context, newsId int) (transaction_detail.Domain, error) {
	rec := TransactionDetail{}
	err := ar.Conn.Where("id = ?", newsId).First(&rec).Error
	if err != nil {
		return transaction_detail.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlTransactionDetailRepository) GetByDescription(ctx context.Context, newsTitle string) (transaction_detail.Domain, error) {
	rec := TransactionDetail{}
	err := ar.Conn.Where("description = ?", newsTitle).First(&rec).Error
	if err != nil {
		return transaction_detail.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlTransactionDetailRepository) Store(ctx context.Context, newsDomain *transaction_detail.Domain) (transaction_detail.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Create(&rec)
	if result.Error != nil {
		return transaction_detail.Domain{}, result.Error
	}
	// ini buat join
	err := ar.Conn.Preload("User").First(&rec, rec.Id).Error
	if err != nil {
		return transaction_detail.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlTransactionDetailRepository) Update(ctx context.Context, newsDomain *transaction_detail.Domain) (transaction_detail.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Save(&rec)
	if result.Error != nil {
		return transaction_detail.Domain{}, result.Error
	}

	err := ar.Conn.Preload("User").First(&rec, rec.Id).Error
	if err != nil {
		return transaction_detail.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlTransactionDetailRepository) Find(ctx context.Context) ([]transaction_detail.Domain, error) {
	rec := []TransactionDetail{}

	query := ar.Conn.Preload("Coffee").Preload("Transaction")

	err := query.Find(&rec).Error
	if err != nil {
		return []transaction_detail.Domain{}, err
	}

	articleDomain := []transaction_detail.Domain{}
	for _, value := range rec {
		articleDomain = append(articleDomain, value.toDomain())
	}

	return articleDomain, nil
}

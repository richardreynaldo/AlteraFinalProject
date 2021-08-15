package transaction_header

import (
	"context"
	"finalProject/business/transaction_header"

	"gorm.io/gorm"
)

type mysqlTransactionHeaderRepository struct {
	Conn *gorm.DB
}

func NewMySQLTransactionHeaderRepository(conn *gorm.DB) transaction_header.Repository {
	return &mysqlTransactionHeaderRepository{
		Conn: conn,
	}
}

func (ar *mysqlTransactionHeaderRepository) Fetch(ctx context.Context, page, perpage int) ([]transaction_header.Domain, int, error) {
	rec := []TransactionHeader{}

	offset := (page - 1) * perpage
	err := ar.Conn.Preload("User").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []transaction_header.Domain{}, 0, err
	}

	var totalData int64
	err = ar.Conn.Count(&totalData).Error
	if err != nil {
		return []transaction_header.Domain{}, 0, err
	}

	var domainNews []transaction_header.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (ar *mysqlTransactionHeaderRepository) GetByID(ctx context.Context, newsId int) (transaction_header.Domain, error) {
	rec := TransactionHeader{}
	err := ar.Conn.Where("id = ?", newsId).First(&rec).Error
	if err != nil {
		return transaction_header.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlTransactionHeaderRepository) GetByDescription(ctx context.Context, newsTitle string) (transaction_header.Domain, error) {
	rec := TransactionHeader{}
	err := ar.Conn.Where("description = ?", newsTitle).First(&rec).Error
	if err != nil {
		return transaction_header.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlTransactionHeaderRepository) Store(ctx context.Context, newsDomain *transaction_header.Domain) (transaction_header.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Create(&rec)
	if result.Error != nil {
		return transaction_header.Domain{}, result.Error
	}
	// ini buat join
	err := ar.Conn.Preload("User").First(&rec, rec.Id).Error
	if err != nil {
		return transaction_header.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlTransactionHeaderRepository) Update(ctx context.Context, newsDomain *transaction_header.Domain) (transaction_header.Domain, error) {
	rec := fromDomain(newsDomain)

	result := ar.Conn.Save(&rec)
	if result.Error != nil {
		return transaction_header.Domain{}, result.Error
	}

	err := ar.Conn.Preload("User").First(&rec, rec.Id).Error
	if err != nil {
		return transaction_header.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlTransactionHeaderRepository) Find(ctx context.Context) ([]transaction_header.Domain, error) {
	rec := []TransactionHeader{}

	query := ar.Conn.Preload("User")

	err := query.Find(&rec).Error
	if err != nil {
		return []transaction_header.Domain{}, err
	}

	articleDomain := []transaction_header.Domain{}
	for _, value := range rec {
		articleDomain = append(articleDomain, value.toDomain())
	}

	return articleDomain, nil
}

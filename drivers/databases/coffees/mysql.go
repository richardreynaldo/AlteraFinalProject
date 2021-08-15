package coffees

import (
	"context"
	"finalProject/business/coffees"

	"gorm.io/gorm"
)

type mysqlcoffeesRepository struct {
	Conn *gorm.DB
}

func NewMySQLCoffeesRepository(conn *gorm.DB) coffees.Repository {
	return &mysqlcoffeesRepository{
		Conn: conn,
	}
}

func (ar *mysqlcoffeesRepository) Fetch(ctx context.Context, page, perpage int) ([]coffees.Domain, int, error) {
	rec := []Coffees{}

	offset := (page - 1) * perpage
	err := ar.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []coffees.Domain{}, 0, err
	}

	var totalData int64
	err = ar.Conn.Count(&totalData).Error
	if err != nil {
		return []coffees.Domain{}, 0, err
	}

	var domainNews []coffees.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (ar *mysqlcoffeesRepository) GetByID(ctx context.Context, coffeesId int) (coffees.Domain, error) {
	rec := Coffees{}
	err := ar.Conn.Where("id = ?", coffeesId).First(&rec).Error
	if err != nil {
		return coffees.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ar *mysqlcoffeesRepository) Store(ctx context.Context, coffeesDomain *coffees.Domain) (coffees.Domain, error) {
	rec := fromDomain(coffeesDomain)

	result := ar.Conn.Create(&rec)
	if result.Error != nil {
		return coffees.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlcoffeesRepository) Update(ctx context.Context, coffeesDomain *coffees.Domain) (coffees.Domain, error) {
	rec := fromDomain(coffeesDomain)

	result := ar.Conn.Save(&rec)
	if result.Error != nil {
		return coffees.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ar *mysqlcoffeesRepository) Find(ctx context.Context) ([]coffees.Domain, error) {
	rec := []Coffees{}

	query := ar.Conn

	err := query.Find(&rec).Error
	if err != nil {
		return []coffees.Domain{}, err
	}

	coffeeDomain := []coffees.Domain{}
	for _, value := range rec {
		coffeeDomain = append(coffeeDomain, value.toDomain())
	}

	return coffeeDomain, nil
}

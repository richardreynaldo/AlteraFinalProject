package transaction_detail

import (
	transactionDetailUsecase "finalProject/business/transaction_detail"
	"finalProject/drivers/databases/coffees"
	"time"
)

type TransactionDetail struct {
	Id            int
	Price         float64
	Quantity      int
	TransactionId int
	CoffeeId      int
	Coffee        coffees.Coffees
	CreatedAt     time.Time `gorm:"<-create"`
	UpdatedAt     time.Time
}

func fromDomain(domain *transactionDetailUsecase.Domain) *TransactionDetail {
	return &TransactionDetail{
		Id:            domain.Id,
		TransactionId: domain.TransactionId,
		CoffeeId:      domain.CoffeeId,
		Price:         domain.Price,
		Quantity:      domain.Quantity,
	}
}

func (rec *TransactionDetail) toDomain() transactionDetailUsecase.Domain {
	return transactionDetailUsecase.Domain{
		Id:            rec.Id,
		Price:         rec.Price,
		Quantity:      rec.Quantity,
		TransactionId: rec.TransactionId,
		CoffeeId:      rec.CoffeeId,
		CoffeeName:    rec.Coffee.Name,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}

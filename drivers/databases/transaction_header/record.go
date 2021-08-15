package transaction_header

import (
	transactionHeaderUsecase "finalProject/business/transaction_header"
	"finalProject/drivers/databases/users"
	"time"
)

type TransactionHeader struct {
	Id            int
	TotalPrice    float64
	TotalQuantity int
	Status        string
	UserId        int
	User          users.User
	CreatedAt     time.Time `gorm:"<-create"`
	UpdatedAt     time.Time
}

func fromDomain(domain *transactionHeaderUsecase.Domain) *TransactionHeader {
	return &TransactionHeader{
		Id:            domain.Id,
		UserId:        domain.UserId,
		TotalPrice:    domain.TotalPrice,
		TotalQuantity: domain.TotalQuantity,
		Status:        domain.Status,
	}
}

func (rec *TransactionHeader) toDomain() transactionHeaderUsecase.Domain {
	return transactionHeaderUsecase.Domain{
		Id:            rec.Id,
		TotalPrice:    rec.TotalPrice,
		TotalQuantity: rec.TotalQuantity,
		Status:        rec.Status,
		UserId:        rec.UserId,
		UserName:      rec.User.Name,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}

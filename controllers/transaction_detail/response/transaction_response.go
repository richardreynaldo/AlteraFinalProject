package response

import (
	transaction_detail "finalProject/business/transaction_detail"
	"time"
)

type TransactionDetail struct {
	Id            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	TransactionId int       `json:"transaction_id"`
	CoffeeId      int       `json:"coffee_id"`
	Price         float64   `json:"price"`
	Quantity      int       `json:"quantity"`
}

func FromDomain(domain transaction_detail.Domain) TransactionDetail {
	return TransactionDetail{
		Id:            domain.Id,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		TransactionId: domain.TransactionId,
		CoffeeId:      domain.CoffeeId,
		Price:         domain.Price,
		Quantity:      domain.Quantity,
	}
}

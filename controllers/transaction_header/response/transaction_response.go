package response

import (
	transaction_header "finalProject/business/transaction_header"
	"time"
)

type TransactionHeader struct {
	Id            int       `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UserId        int       `json:"user_id"`
	TotalPrice    float64   `json:"total_price"`
	TotalQuantity int       `json:"total_quantity"`
	Status        string    `json:"status"`
}

func FromDomain(domain transaction_header.Domain) TransactionHeader {
	return TransactionHeader{
		Id:            domain.Id,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		UserId:        domain.UserId,
		TotalPrice:    domain.TotalPrice,
		TotalQuantity: domain.TotalQuantity,
		Status:        domain.Status,
	}
}

package request

import (
	"finalProject/business/transaction_detail"
)

type CreateTransactionDetail struct {
	TransactionId int     `json:"transaction_id"`
	CoffeeId      int     `json:"coffee_id"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
}

func (req *CreateTransactionDetail) ToDomain() *transaction_detail.Domain {
	return &transaction_detail.Domain{
		TransactionId: req.TransactionId,
		Price:         req.Price,
		Quantity:      req.Quantity,
		CoffeeId:      req.CoffeeId,
	}
}

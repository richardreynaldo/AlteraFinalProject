package request

import (
	"finalProject/business/transaction_header"
	"finalProject/controllers/transaction_detail/request"
)

type CreateTransactionHeader struct {
	UserId            int                               `json:"user_id"`
	TotalPrice        float64                           `json:"total_price"`
	TotalQuantity     int                               `json:"total_quantity"`
	Status            string                            `json:"status"`
	TransactionDetail []request.CreateTransactionDetail `json:"detail"`
}

func (req *CreateTransactionHeader) ToDomain() *transaction_header.Domain {
	return &transaction_header.Domain{
		UserId:        req.UserId,
		TotalPrice:    req.TotalPrice,
		TotalQuantity: req.TotalQuantity,
		Status:        req.Status,
		
	}
}

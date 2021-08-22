package transaction_header

import (
	"context"
	"finalProject/business"
	"finalProject/business/transaction_detail"
	"time"
)

type transactionHeaderUsecase struct {
	transactionHeaderRepository Repository
	transactionDetailUsecase    transaction_detail.Usecase
	contextTimeout              time.Duration
}

func NewTransactionHeaderUsecase(cr Repository, uc transaction_detail.Usecase, timeout time.Duration) Usecase {
	return &transactionHeaderUsecase{
		transactionHeaderRepository: cr,
		contextTimeout:              timeout,
		transactionDetailUsecase:    uc,
	}
}

func (cu *transactionHeaderUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := cu.transactionHeaderRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (cu *transactionHeaderUsecase) GetByID(ctx context.Context, transactionHeaderId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if transactionHeaderId <= 0 {
		return Domain{}, business.ErrNewsIDResource
	}
	res, err := cu.transactionHeaderRepository.GetByID(ctx, transactionHeaderId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *transactionHeaderUsecase) Store(ctx context.Context, transactionHeaderDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	result, err := cu.transactionHeaderRepository.Store(ctx, transactionHeaderDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (cu *transactionHeaderUsecase) Update(ctx context.Context, transactionHeaderDomain *Domain) (*Domain, error) {
	existingArticle, err := cu.transactionHeaderRepository.GetByID(ctx, transactionHeaderDomain.Id)
	if err != nil {
		return &Domain{}, err
	}
	transactionHeaderDomain.Id = existingArticle.Id

	result, err := cu.transactionHeaderRepository.Update(ctx, transactionHeaderDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *transactionHeaderUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.transactionHeaderRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

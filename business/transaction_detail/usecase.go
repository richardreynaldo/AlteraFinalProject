package transaction_detail

import (
	"context"
	"finalProject/business"
	"time"
)

type transactionDetailUsecase struct {
	transactionDetailRepository Repository
	contextTimeout              time.Duration
}

func NewTransactionDetailUsecase(cr Repository, timeout time.Duration) Usecase {
	return &transactionDetailUsecase{
		transactionDetailRepository: cr,
		contextTimeout:              timeout,
	}
}

func (cu *transactionDetailUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := cu.transactionDetailRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (cu *transactionDetailUsecase) GetByID(ctx context.Context, transactionHeaderId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if transactionHeaderId <= 0 {
		return Domain{}, business.ErrNewsIDResource
	}
	res, err := cu.transactionDetailRepository.GetByID(ctx, transactionHeaderId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *transactionDetailUsecase) Store(ctx context.Context, transactionHeaderDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	result, err := cu.transactionDetailRepository.Store(ctx, transactionHeaderDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (cu *transactionDetailUsecase) Update(ctx context.Context, transactionHeaderDomain *Domain) (*Domain, error) {
	existingArticle, err := cu.transactionDetailRepository.GetByID(ctx, transactionHeaderDomain.Id)
	if err != nil {
		return &Domain{}, err
	}
	transactionHeaderDomain.Id = existingArticle.Id

	result, err := cu.transactionDetailRepository.Update(ctx, transactionHeaderDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *transactionDetailUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.transactionDetailRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

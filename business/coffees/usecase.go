package coffees

import (
	"context"
	"finalProject/business"
	"time"
)

type coffeesUsecase struct {
	coffeesRepository Repository
	contextTimeout    time.Duration
}

func NewCoffeesUsecase(cr Repository, timeout time.Duration) Usecase {
	return &coffeesUsecase{
		coffeesRepository: cr,
		contextTimeout:    timeout,
	}
}

func (cu *coffeesUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := cu.coffeesRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (cu *coffeesUsecase) GetByID(ctx context.Context, articlesId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if articlesId <= 0 {
		return Domain{}, business.ErrNewsIDResource
	}
	res, err := cu.coffeesRepository.GetByID(ctx, articlesId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *coffeesUsecase) Store(ctx context.Context, coffeeDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	result, err := cu.coffeesRepository.Store(ctx, coffeeDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (cu *coffeesUsecase) Update(ctx context.Context, articleDomain *Domain) (*Domain, error) {
	existingArticle, err := cu.coffeesRepository.GetByID(ctx, articleDomain.Id)
	if err != nil {
		return &Domain{}, err
	}
	articleDomain.Id = existingArticle.Id

	result, err := cu.coffeesRepository.Update(ctx, articleDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *coffeesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.coffeesRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

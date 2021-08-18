package reviews

import (
	"context"
	"finalProject/business"
	"finalProject/business/users"
	"strings"
	"time"
)

type reviewUsecase struct {
	reviewRepository Repository
	userUsecase      users.UseCase
	contextTimeout   time.Duration
}

func NewArticleUsecase(ar Repository, uc users.UseCase, timeout time.Duration) Usecase {
	return &reviewUsecase{
		reviewRepository: ar,
		userUsecase:      uc,
		contextTimeout:   timeout,
	}
}

func (au *reviewUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := au.reviewRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (au *reviewUsecase) GetByID(ctx context.Context, articlesId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if articlesId <= 0 {
		return Domain{}, business.ErrNewsIDResource
	}
	res, err := au.reviewRepository.GetByID(ctx, articlesId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (au *reviewUsecase) GetByDescription(ctx context.Context, articleDescription string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if strings.TrimSpace(articleDescription) == "" {
		return Domain{}, business.ErrNewsTitleResource
	}
	res, err := au.reviewRepository.GetByDescription(ctx, articleDescription)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (au *reviewUsecase) Store(ctx context.Context, articleDomain *Domain, userId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	_, err := au.userUsecase.GetByID(ctx, userId)
	if err != nil {
		return Domain{}, business.ErrCategoryNotFound
	}
	articleDomain.UserId = userId

	result, err := au.reviewRepository.Store(ctx, articleDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (au *reviewUsecase) Update(ctx context.Context, articleDomain *Domain) (*Domain, error) {
	existingArticle, err := au.reviewRepository.GetByID(ctx, articleDomain.Id)
	if err != nil {
		return &Domain{}, err
	}
	articleDomain.Id = existingArticle.Id

	result, err := au.reviewRepository.Update(ctx, articleDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *reviewUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.reviewRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

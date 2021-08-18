package articles

import (
	"context"
	"finalProject/business"
	"finalProject/business/users"
	"strings"
	"time"
)

type articlesUsecase struct {
	articlesRepository Repository
	userUsecase        users.UseCase
	contextTimeout     time.Duration
}

func NewArticleUsecase(ar Repository, uc users.UseCase, timeout time.Duration) Usecase {
	return &articlesUsecase{
		articlesRepository: ar,
		userUsecase:        uc,
		contextTimeout:     timeout,
	}
}

func (au *articlesUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := au.articlesRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (au *articlesUsecase) GetByID(ctx context.Context, articlesId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if articlesId <= 0 {
		return Domain{}, business.ErrNewsIDResource
	}
	res, err := au.articlesRepository.GetByID(ctx, articlesId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (au *articlesUsecase) GetByDescription(ctx context.Context, articleDescription string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if strings.TrimSpace(articleDescription) == "" {
		return Domain{}, business.ErrNewsTitleResource
	}
	res, err := au.articlesRepository.GetByDescription(ctx, articleDescription)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (au *articlesUsecase) Store(ctx context.Context, articleDomain *Domain, userId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	_, err := au.userUsecase.GetByID(ctx, userId)
	if err != nil {
		return Domain{}, business.ErrCategoryNotFound
	}

	existedNews, err := au.articlesRepository.GetByDescription(ctx, articleDomain.Description)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if existedNews != (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}
	articleDomain.UserId = userId

	result, err := au.articlesRepository.Store(ctx, articleDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (au *articlesUsecase) Update(ctx context.Context, articleDomain *Domain) (*Domain, error) {
	existingArticle, err := au.articlesRepository.GetByID(ctx, articleDomain.Id)
	if err != nil {
		return &Domain{}, err
	}
	articleDomain.Id = existingArticle.Id

	result, err := au.articlesRepository.Update(ctx, articleDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *articlesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.articlesRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

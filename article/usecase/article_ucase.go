package usecase

import (
	"backend/domain"
	"context"
	"time"
)

type articleUsecase struct {
	articleRepo    domain.ArticleRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(a domain.ArticleRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) Fetch(c context.Context, limit, offset int) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.Fetch(ctx, limit, offset)
}

func (a *articleUsecase) GetByID(c context.Context, id int) (domain.Article, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.GetByID(ctx, id)
}

func (a *articleUsecase) Store(c context.Context, m *domain.Article) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.Store(ctx, m)
}

func (a *articleUsecase) Update(c context.Context, m *domain.Article) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.Update(ctx, m)
}

func (a *articleUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.Delete(ctx, id)
}

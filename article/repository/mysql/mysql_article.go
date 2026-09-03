package mysql

import (
	"backend/domain"
	"context"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(db *gorm.DB) domain.ArticleRepository {
	return &mysqlArticleRepository{DB: db}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context, limit, offset int) (res []domain.Article, err error) {
	err = m.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&res).Error
	return
}

func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int) (res domain.Article, err error) {
	err = m.DB.WithContext(ctx).First(&res, id).Error
	return
}

func (m *mysqlArticleRepository) Store(ctx context.Context, a *domain.Article) (err error) {
	err = m.DB.WithContext(ctx).Create(a).Error
	return
}

func (m *mysqlArticleRepository) Update(ctx context.Context, a *domain.Article) (err error) {
	err = m.DB.WithContext(ctx).Save(a).Error
	return
}

func (m *mysqlArticleRepository) Delete(ctx context.Context, id int) (err error) {
	err = m.DB.WithContext(ctx).Delete(&domain.Article{}, id).Error
	return
}

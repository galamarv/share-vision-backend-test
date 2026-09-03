package mysql

import (
	"backend/domain"
	"context"
	"time"

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

func (m *mysqlArticleRepository) Update(ctx context.Context, ar *domain.Article) error {
	// Use .Updates to only update specific fields, avoiding overwriting created_date
	err := m.DB.Model(&domain.Article{}).Where("id = ?", ar.ID).Updates(map[string]interface{}{
		"title":        ar.Title,
		"content":      ar.Content,
		"category":     ar.Category,
		"status":       ar.Status,
		"updated_date": time.Now(),
	}).Error

	return err
}

func (m *mysqlArticleRepository) Delete(ctx context.Context, id int) (err error) {
	err = m.DB.WithContext(ctx).Delete(&domain.Article{}, id).Error
	return
}

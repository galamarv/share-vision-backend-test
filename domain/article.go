package domain

import (
	"context"
	"time"
)

type Article struct {
	ID          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title       string    `json:"title" validate:"required,min=20" gorm:"column:title;type:varchar(200);not null"`
	Content     string    `json:"content" validate:"required,min=200" gorm:"column:content;type:text;not null"`
	Category    string    `json:"category" validate:"required,min=3" gorm:"column:category;type:varchar(100);not null"`
	Status      string    `json:"status" validate:"required,oneof=Publish Draft Thrash publish draft thrash" gorm:"column:status;type:varchar(100);not null"`
	CreatedDate time.Time `json:"created_date" gorm:"column:created_date;type:timestamp;autoCreateTime"`
	UpdatedDate time.Time `json:"updated_date" gorm:"column:updated_date;type:timestamp;autoUpdateTime"`
}

func (Article) TableName() string {
	return "posts"
}

type ArticleRepository interface {
	Fetch(ctx context.Context, limit, offset int) ([]Article, error)
	GetByID(ctx context.Context, id int) (Article, error)
	Store(ctx context.Context, a *Article) error
	Update(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int) error
}

type ArticleUsecase interface {
	Fetch(ctx context.Context, limit, offset int) ([]Article, error)
	GetByID(ctx context.Context, id int) (Article, error)
	Store(ctx context.Context, a *Article) error
	Update(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int) error
}

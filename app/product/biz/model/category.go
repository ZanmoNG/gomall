package model

import (
	"golang.org/x/net/context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (category []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&category).Error
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;index;"`
	Qty       uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserID: item.UserID, ProductId: item.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserID: item.UserID, ProductId: item.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	}
	return db.WithContext(ctx).Create(item).Error
}

package auth

import (
	"context"

	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}
type Interface interface {
	DBExampleGet() error
}

func NewAuthHandler(ctx context.Context, db *gorm.DB) (Interface, error) {
	return &AuthHandler{db: db}, nil
}

func (h *AuthHandler) DBExampleGet() error {
	tx := h.db.Begin()
	isCommitted := false
	defer func() {
		if !isCommitted {
			tx.Rollback()
		}
	}()

	//tx.Model(&model.Test{}).Create(&model.Test{ID: 1})

	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	isCommitted = true
	return tx.Error
}

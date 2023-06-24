package model

type Test struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	AccountID string `gorm:"unique"`
}

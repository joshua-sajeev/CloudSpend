package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"index"`
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

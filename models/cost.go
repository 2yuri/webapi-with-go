package models

import (
	"time"

	"gorm.io/gorm"
)

type Cost struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	IDCostCategory int            `json:"id_cost_category"`
	IDProvider     int            `json:"id_provider"`
	PurchaseDate   time.Time      `json:"purchase_date"`
	DueDate        time.Time      `json:"due_date"`
	Total          float32        `json:"total"`
	CreatedAt      time.Time      `json:"created"`
	UpdatedAt      time.Time      `json:"updated"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted"`
}

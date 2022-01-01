package models

import (
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Cnpj      int            `json:"cnpj"`
	Telephone int            `json:"telephone"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

package models

type CostCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

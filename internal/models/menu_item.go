package models

type MenuItem struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}
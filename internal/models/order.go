package models

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items"`
}

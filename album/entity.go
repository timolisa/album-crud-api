package album

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	ID     int     `gorm:"primaryKey"`
	Title  string  `gorm:"size:255;not null" json:"title"`
	Artist string  `gorm:"size:255;not null" json:"artist"`
	Price  float64 `json:"price"`
}
package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TotalPrice          int32 `gorm:"not null; default:null" json:"total_price"`
	TransactionDetailID uint8 `json:"td_id"`
}

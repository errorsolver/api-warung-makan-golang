package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	UserID      int         `json:"user_id"`
	ProductID   int         `json:"product_id"`
	Amount      int         `gorm:"type:int; not null" json:"amount"`
	Transaction Transaction `gorm:"ForeignKey:TransactionDetailID"`
}

// type TransactionDetailStruct struct {
// 	gorm.Model
// 	UserID      int         `json:"user_id"`
// 	ProductID   int         `json:"product_id"`
// 	Amount      int         `gorm:"type:int; not null" json:"amount"`
// }

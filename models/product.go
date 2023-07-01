package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName        string              `gorm:"type:varchar(10); unique; not null; default:null" json:"product_name"`
	Description        string              `gorm:"type:text; not null" json:"description"`
	Price              int                 `gorm:"type:int; not null" json:"price"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:ProductID"`
}

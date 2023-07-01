package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username           string              `gorm:"type:varchar(10); unique; not null" json:"username"`
	Password           string              `gorm:"type:varchar(10); not null" json:"password"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:UserID"`
}

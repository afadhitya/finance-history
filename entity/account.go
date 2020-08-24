package entity

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Name   string      `json:"name"`
	Type   AccountType `json:"account_type"`
	IDType uint        `json:"id_type"`
}

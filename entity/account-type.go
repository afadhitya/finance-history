package entity

import (
	"github.com/jinzhu/gorm"
)

type AccountType struct {
	gorm.Model
	TypeName string `json: "type_name"` 
}
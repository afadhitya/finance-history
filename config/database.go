package config

import (
	"github.com/afadhitya/finance-history/entity"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	dbConnectionStr := "root@/keuangan-db?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open("mysql", dbConnectionStr)
	if err != nil {
		panic("Failed to connect to DB")
	}

	DB.DropTableIfExists(
		&entity.Account{},
		&entity.AccountType{},
	)
	DB.AutoMigrate(
		&entity.Account{},
		&entity.AccountType{},
	)

}

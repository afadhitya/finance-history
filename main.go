package main

import (
	"fmt"

	"github.com/afadhitya/finance-history/config"
	"github.com/afadhitya/finance-history/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	fmt.Println("Apps started")
	setRouter()
}

func setRouter() {
	router := gin.Default()
	accountType := router.Group("api/v1/account-type")
	{
		accountType.POST("/", service.SaveAccountType)
	}

	router.Use(cors.Default())
	router.Run(":8081")
}

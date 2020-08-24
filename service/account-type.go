package service

import (
	"net/http"

	"github.com/afadhitya/finance-history/config"
	"github.com/afadhitya/finance-history/entity"
	"github.com/gin-gonic/gin"
)

func SaveAccountType(c *gin.Context) {
	accountType := getAttribute(c)
	config.DB.Save(&accountType)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Account Type has saved to DB",
		"resourceId": accountType.ID,
	})
}

func getAttribute(c *gin.Context) entity.AccountType {
	accountType := entity.AccountType{
		TypeName: c.PostForm("account_type_name"),
	}
	return accountType
}

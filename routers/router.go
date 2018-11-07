package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mikejsoh/captrade/controllers"
)

var router *gin.Engine

// Init Router
func Init() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		companyGroup := v1.Group("company")
		{
			company := new(controllers.CompanyController)
			companyGroup.POST("/", company.CreateCompany)
			companyGroup.GET("/", company.FetchAllCompany)
			companyGroup.GET("/:id", company.FetchSingleCompany)
			companyGroup.PUT("/:id", company.UpdateCompany)
			companyGroup.DELETE("/:id", company.DeleteCompany)
		}

		financeGroup := v1.Group("finance")
		{
			finance := new(controllers.FinanceController)
			financeGroup.POST("/", finance.CreateFinance)
			// financeGroup.GET("/", finance.FetchAllFinance)
			// financeGroup.GET("/:id", finance.FetchSingleFinance)
			// financeGroup.PUT("/:id", finance.UpdateFinance)
			// financeGroup.DELETE("/:id", finance.DeleteFinance)
		}
	}
	router.Run(":8080")
}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikejsoh/captrade/db"
	"github.com/mikejsoh/captrade/models"
)

// FinanceController struct
type FinanceController struct{}

// CreateFinance creates a new financial record
func (h FinanceController) CreateFinance(c *gin.Context) {
	year, _ := strconv.Atoi(c.PostForm("year"))
	companyID, _ := strconv.Atoi(c.PostForm("companyId"))
	productionLevel, _ := strconv.Atoi(c.PostForm("productionLevel"))
	salesPrice, _ := strconv.ParseFloat(c.PostForm("salesPrice"), 64)
	unitCost, _ := strconv.ParseFloat(c.PostForm("unitCost"), 64)
	totalProductionCost, _ := strconv.ParseFloat(c.PostForm("totalProdCost"), 64)
	grossMargin, _ := strconv.ParseFloat(c.PostForm("grossMargin"), 64)
	fixedCost, _ := strconv.ParseFloat(c.PostForm("fixedCost"), 64)
	profit, _ := strconv.ParseFloat(c.PostForm("profit"), 64)
	permitID, _ := strconv.Atoi(c.PostForm("permitId"))
	carbonMarketID, _ := strconv.Atoi(c.PostForm("carbonMarketId"))

	finance := models.Finance{
		CompanyID:           companyID,
		Year:                year,
		ProductionLevel:     productionLevel,
		SalesPrice:          salesPrice,
		UnitCost:            unitCost,
		TotalProductionCost: totalProductionCost,
		GrossMargin:         grossMargin,
		FixedCost:           fixedCost,
		Profit:              profit,
		PermitID:            permitID,
		CarbonMarketID:      carbonMarketID,
	}
	db := db.GetDB()
	db.Save(&finance)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Finance item created successfully!",
		"resourceId": finance.FinancialID,
	})

}

// FetchAllFinance fetches all financial records
func (h FinanceController) FetchAllFinance(c *gin.Context) {
	var finances []models.Finance
	db := db.GetDB()

	companyID := c.Request.URL.Query().Get("company_id")
	fmt.Println("Company ID: ", companyID)

	if companyID != "" {
		db.Where("company_id = ?", companyID).Find(&finances)
	} else {
		db.Find(&finances)
	}

	if len(finances) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No finances found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": finances})
}

// FetchSingleFinance fetches a single financial record based on financial id
func (h FinanceController) FetchSingleFinance(c *gin.Context) {

}

// UpdateFinance updates a single financial record based on financial id
func (h FinanceController) UpdateFinance(c *gin.Context) {

}

// DeleteFinance deletes financial record based on financial id
func (h FinanceController) DeleteFinance(c *gin.Context) {

}

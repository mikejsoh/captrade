package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikejsoh/captrade/db"
	"github.com/mikejsoh/captrade/models"
)

// CarbonMarketController struct
type CarbonMarketController struct{}

// CreateCarbonMarket creates a carbon market
func (h CarbonMarketController) CreateCarbonMarket(c *gin.Context) {
	carbonIntensity, _ := strconv.ParseFloat(c.PostForm("carbonIntensity"), 64)
	carbonProduction, _ := strconv.Atoi(c.PostForm("carbonProduction"))
	carbonPermitAllowance, _ := strconv.Atoi(c.PostForm("carbonPermitAllowance"))
	carbonPermitSellMax, _ := strconv.Atoi(c.PostForm("carbonPermitSellMax"))

	carbonMarket := models.CarbonMarket{
		CarbonIntensity:       carbonIntensity,
		CarbonProduction:      carbonProduction,
		CarbonPermitAllowance: carbonPermitAllowance,
		CarbonPermitSellMax:   carbonPermitSellMax,
	}
	db := db.GetDB()
	db.Save(&carbonMarket)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Carbon Market item created successfully!",
		"resourceId": carbonMarket.ID,
	})
}

// FetchAllCarbonMarket fetches all carbon markets
func (h CarbonMarketController) FetchAllCarbonMarket(c *gin.Context) {
	var carbonmarkets []models.CarbonMarket
	db := db.GetDB()
	db.Find(&carbonmarkets)

	if len(carbonmarkets) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No carbon markets found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carbonmarkets})
}

// FetchSingleCarbonMarket fetches a single carbon market based on id
func (h CarbonMarketController) FetchSingleCarbonMarket(c *gin.Context) {
	var carbonMarket models.CarbonMarket
	carbonMarketID := c.Param("id")

	db := db.GetDB()
	db.First(&carbonMarket, carbonMarketID)

	if carbonMarket.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No carbon market found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carbonMarket})
}

// UpdateCarbonMarket updates a single carbon market based on id
func (h CarbonMarketController) UpdateCarbonMarket(c *gin.Context) {
	var carbonMarket models.CarbonMarket
	carbonMarketID := c.Param("id")

	carbonIntensity, _ := strconv.ParseFloat(c.PostForm("carbonIntensity"), 64)
	carbonProduction, _ := strconv.Atoi(c.PostForm("carbonProduction"))
	carbonPermitAllowance, _ := strconv.Atoi(c.PostForm("carbonPermitAllowance"))
	carbonPermitSellMax, _ := strconv.Atoi(c.PostForm("carbonPermitSellMax"))

	db := db.GetDB()
	db.First(&carbonMarket, carbonMarketID)

	if carbonMarket.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No carbon market found!"})
		return
	}

	// http://gorm.io/docs/update.html
	db.Model(&carbonMarket).Updates(models.CarbonMarket{
		CarbonIntensity:       carbonIntensity,
		CarbonProduction:      carbonProduction,
		CarbonPermitAllowance: carbonPermitAllowance,
		CarbonPermitSellMax:   carbonPermitSellMax,
	})

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carbonMarket})

}

// DeleteCarbonMarket removes a carbon market based on id
func (h CarbonMarketController) DeleteCarbonMarket(c *gin.Context) {
	var carbonMarket models.CarbonMarket
	carbonMarketID := c.Param("id")

	db := db.GetDB()
	db.First(&carbonMarket, carbonMarketID)

	if carbonMarket.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No carbon market found!"})
		return
	}

	db.Delete(&carbonMarket)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Carbon Market deleted successfully!"})
}

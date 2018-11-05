package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikejsoh/captrade/db"
	"github.com/mikejsoh/captrade/models"
)

type CompanyController struct{}

func (h CompanyController) CreateCompany(c *gin.Context) {
	company := models.Company{Name: c.PostForm("name")}
	db := db.GetDB()
	db.Save(&company)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Company item created successfully!",
		"resourceId": company.ID,
	})
}

func (h CompanyController) FetchAllCompany(c *gin.Context) {
	var companies []models.Company
	db := db.GetDB()
	db.Find(&companies)

	if len(companies) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": companies})
}

func (h CompanyController) FetchSingleCompany(c *gin.Context) {
	var company models.Company
	companyID := c.Param("id")

	db := db.GetDB()
	db.First(&company, companyID)

	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": company})
}

func (h CompanyController) UpdateCompany(c *gin.Context) {

}

func (h CompanyController) DeleteCompany(c *gin.Context) {

}

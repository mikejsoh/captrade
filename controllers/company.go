package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikejsoh/captrade/db"
	"github.com/mikejsoh/captrade/models"
)

// CompanyController struct
type CompanyController struct{}

// CreateCompany creates a new company
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

// FetchAllCompany fetches all companies
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

// FetchSingleCompany fetches a single company based on id
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

// UpdateCompany updates a company specified by id
func (h CompanyController) UpdateCompany(c *gin.Context) {
	var company models.Company
	companyID := c.Param("id")

	db := db.GetDB()
	db.First(&company, companyID)

	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&company).Update("Name", c.PostForm("name"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Company updated successfully!"})
}

// DeleteCompany delete company record by id
func (h CompanyController) DeleteCompany(c *gin.Context) {
	var company models.Company
	companyID := c.Param("id")

	db := db.GetDB()
	db.First(&company, companyID)

	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&company)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

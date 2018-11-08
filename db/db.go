package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // needed for now
	"github.com/mikejsoh/captrade/models"
)

var db *gorm.DB

// Init DB connection
func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/captrade")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&models.Company{}, &models.CarbonMarket{}, &models.Finance{})

	// Add Foreign Keys Manually
	// Issue: https://github.com/jinzhu/gorm/issues/450
	db.Model(&models.Finance{}).AddForeignKey("company_id", "companies(company_id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Finance{}).AddForeignKey("carbonmarket_id", "carbon_markets(id)", "RESTRICT", "RESTRICT")
}

// GetDB returns db
func GetDB() *gorm.DB {
	return db
}

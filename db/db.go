package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db.AutoMigrate(&models.Company{}, &models.Finance{})

	// Add Foreign Keys Manually
	// Issue: https://github.com/jinzhu/gorm/issues/450
	db.Model(&models.Finance{}).AddForeignKey("company_id", "companies(company_id)", "RESTRICT", "RESTRICT")
}

// GetDB returns db
func GetDB() *gorm.DB {
	return db
}

package models

type (
	// Company describes a Company type
	Company struct {
		ID   int    `gorm:"column:company_id;primary_key" json:"id"`
		Name string `gorm:"column:company_name;not null" json:"name"`
	}
)

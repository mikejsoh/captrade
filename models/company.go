package models

type (
	// Company describes a Company type
	Company struct {
		ID       int       `gorm:"column:company_id;primary_key" json:"id"`
		Name     string    `gorm:"column:company_name;not null" json:"name" sql:"type:varchar(55)"`
		Finances []Finance `gorm:"foreignkey:ID; association_foreignkey:CompanyID"`
	}
)

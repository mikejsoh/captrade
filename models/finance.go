package models

type (
	// Finance struct describes a Finance type
	Finance struct {
		FinancialID         int     `gorm:"column:finance_id;primary_key" json:"id"`
		CompanyID           int     `gorm:"column:company_id;not null;foreignkey:ID" json:"companyId"`
		Year                int     `gorm:"column:year" json:"year"`
		ProductionLevel     int     `gorm:"column:production_level" json:"productionLevel"`
		SalesPrice          float64 `gorm:"column:sales_price" json:"salesPrice" sql:"type:decimal(13,2)"`
		UnitCost            float64 `gorm:"column:unit_cost" json:"unitCost" sql:"type:decimal(13,2)"`
		TotalProductionCost float64 `gorm:"column:total_prod_cost" json:"totalProdCost" sql:"type:decimal(13,2)"`
		GrossMargin         float64 `gorm:"column:gross_margin" json:"grossMargin" sql:"type:decimal(13,2)"`
		FixedCost           float64 `gorm:"column:fixed_cost" json:"fixedCost" sql:"type:decimal(13,2)"`
		Profit              float64 `gorm:"column:profit" json:"profit" sql:"type:decimal(13,2)"`
		PermitID            int     `gorm:"column:permit_id" json:"permitId"`
		CarbonMarketID      int     `gorm:"column:carbonmarket_id;not null;foreignkey:ID" json:"carbonMarketId"`
	}
)

package models

type (
	// CarbonMarket describes a CarbonMarket type
	CarbonMarket struct {
		ID                    int     `gorm:"column:id;primary_key" json:"id"`
		CarbonIntensity       float64 `gorm:"column:carbon_intensity;" json:"carbonIntensity" sql:"type:decimal(4,2)"`
		CarbonProduction      int     `gorm:"column:carbon_production;" json:"carbonProduction"`
		CarbonPermitAllowance int     `gorm:"column:carbon_permit_allowance;" json:"carbonPermitAllowance"`
		CarbonPermitSellMax   int     `gorm:"column:carbon_permit_sell_max;" json:"carbonPermitSellMax"`

		Finances []Finance `gorm:"foreignkey:ID; association_foreignkey:CarbonMarketID"`
	}
)

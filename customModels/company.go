package custommodels

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name      string `json:"name"`
	Location  string `json:"location"`
	CompanyID int    `json:"companyId"`
	Jobs      []*Job `json:"jobs,omitempty"`
}

type NewCompany struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	CompanyID int    `json:"companyId"`
}

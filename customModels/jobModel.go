package custommodels

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyID   int    `json:"companyId"`
}

type NewJob struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyID   int    `json:"companyId"`
}

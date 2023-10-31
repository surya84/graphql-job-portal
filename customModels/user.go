package custommodels

import "gorm.io/gorm"

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type Users struct {
	gorm.Model
	//ID       int    `json:"id" gorm:"primarykey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"Password"`
}

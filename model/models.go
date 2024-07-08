package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

// type Books struct {
// 	gorm.Model
// 	Name  string
// 	Price string
// 	Author
// }

// type Author struct {
// 	Author      string
// 	Publication string
// }

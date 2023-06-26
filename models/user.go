package models

import (
	"github.com/nqnlong1506/user-authentication/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username;not null;unique"`
	Password string `gorm:"password;not null"`
}

func InitializeUser() {
	database.DB.AutoMigrate(&User{})
}

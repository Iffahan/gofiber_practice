// models/user.go
package models

import "gorm.io/gorm"

// Define the User model
type User struct {
	gorm.Model
	Name  string `gorm:"size:100"`
	Email string `gorm:"unique"`
}

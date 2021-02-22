package models

import "gorm.io/gorm"

// User Model
type User struct {
	gorm.Model
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	HashedPassword string `json:"-"`
}

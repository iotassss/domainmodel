package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID      string `gorm:"type:char(36);unique;not null"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
}

type Credential struct {
	gorm.Model
	UserUUID     string `gorm:"type:char(36);unique;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	// TODO: Saltを追加する
}

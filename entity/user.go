package entity

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username       string
	HashedPassword string
	Salt           string
}

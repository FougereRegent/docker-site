package entity

import (
	"docker-site/dto/user"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	HashedPassword string
	Salt           string
	Connection     []UserConnectionModel `gorm:"foreignKey:UserRefer"`
	UserDetails    UserDetailsModel      `gorm:"foreignKey:UserRefer"`
}

type UserConnectionModel struct {
	gorm.Model
	ConnectionDate    time.Time
	DisconnectionDate time.Time
	UserRefer         uint64
}

type UserDetailsModel struct {
	gorm.Model
	FirstName string
	LastName  string
	UserRefer uint64
}

func (o *UserModel) UserModelToUserFrontDTO() user.UserFrontDTO {
	if len(o.Connection) >= 1 {
		return user.UserFrontDTO{
			Id:             uint64(o.ID),
			Name:           o.Username,
			LastConnection: &o.Connection[len(o.Connection)-1].ConnectionDate,
		}
	} else {
		return user.UserFrontDTO{
			Id:             uint64(o.ID),
			Name:           o.Username,
			LastConnection: nil,
		}
	}
}

func (o *UserModel) UserModelToUserDetailsDTO() user.UserDetailsDTO {
	return user.UserDetailsDTO{
		Id:        uint(o.ID),
		Username:  o.Username,
		FirstName: o.UserDetails.FirstName,
		LastName:  o.UserDetails.LastName,
	}
}

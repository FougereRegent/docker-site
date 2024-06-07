package userchecker

import (
	"docker-site/dto/user"
	"docker-site/entity"
	"docker-site/helper"

	"gorm.io/gorm"
)

type SaveUser struct {
	baseComponent
	Db *gorm.DB
}

func (o *SaveUser) Execute(user user.UserCreationDTO) (bool, error) {
	db := o.Db

	hahsedPassword := helper.HashPassword(user.Password)

	userModel := entity.UserModel{
		Username:       user.Name,
		HashedPassword: hahsedPassword.Digest,
		Salt:           hahsedPassword.Salt,
	}

	if result := db.Create(&userModel); result.Error != nil || result.RowsAffected <= 0 {
		return false, result.Error
	}

	if o.Next != nil {
		return o.Next.Execute(user)
	}
	return true, nil
}

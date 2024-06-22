package updatepassword

import (
	"docker-site/dto/user"
	"docker-site/entity"
	"docker-site/helper"
	"errors"

	"gorm.io/gorm"
)

type UpdatePassword struct {
	baseComponent
	Db gorm.DB
}

func (o *UpdatePassword) Execute(id uint, user user.UserUpdatePasswordDTO) (bool, error) {
	var userModel entity.UserModel
	db := o.Db

	if res := db.Find(&userModel, "id = ?", id); res.Error != nil {
		return false, res.Error
	} else if res.RowsAffected < 1 {
		return false, errors.New("User Not found")
	}

	hashedPassord := helper.HashPassword(user.NewPassword)
	userModel.Salt = hashedPassord.Salt
	userModel.HashedPassword = hashedPassord.Digest

	if res := db.Save(&userModel); res.Error != nil {
		return false, res.Error
	}

	if o.Next != nil {
		return o.Next.Execute(id, user)
	}

	return true, nil
}

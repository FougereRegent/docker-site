package updatepassword

import (
	"docker-site/dto/user"
	"docker-site/entity"
	"docker-site/helper"
	"errors"

	"gorm.io/gorm"
)

type CheckOldPassword struct {
	baseComponent
	Db gorm.DB
}

func (o *CheckOldPassword) Execute(id uint, userUpdatePassword user.UserUpdatePasswordDTO) (bool, error) {
	var user entity.UserModel
	db := o.Db

	if ctx := db.First(&user, "id = ?", id); ctx.Error != nil {
		return false, ctx.Error
	} else if ctx.RowsAffected < 1 {
		return false, errors.New("User not found")
	}

	hashedPassword := helper.HashedPassword{
		Digest: user.HashedPassword,
		Salt:   user.Salt,
	}

	if !helper.CheckPassword(userUpdatePassword.OldPassword, &hashedPassword) {
		return false, errors.New("Bad password")
	}

	if o.Next != nil {
		return o.Next.Execute(id, userUpdatePassword)
	}

	return true, nil
}

package userchecker

import (
	"docker-site/dto/user"
	"docker-site/entity"
	"errors"

	"gorm.io/gorm"
)

type CheckIfExist struct {
	baseComponent
	Db *gorm.DB
}

func (o *CheckIfExist) Execute(user user.UserCreationDTO) (bool, error) {
	db := o.Db
	if result := db.Find(&entity.UserModel{}, "username = ?", user.Name); result.RowsAffected >= 1 {
		return false, errors.New("")
	}

	if o.Next != nil {
		return o.Next.Execute(user)
	}

	return true, nil
}

package service

import (
	"docker-site/dto/user"
	"docker-site/entity"
	"docker-site/helper"
	updatepassword "docker-site/service/update_password"
	. "docker-site/service/user_checker"
	"errors"
	"time"

	"gorm.io/gorm"
)

type IUserService interface {
	GetAllUsers() ([]user.UserFrontDTO, error)
	SetConnection(username, password string) error
	CreateUser(user user.UserCreationDTO) error
	DeleteUser(id uint) error
	GetUserDetails(id uint) (*user.UserDetailsDTO, error)
	UpdateUser(id uint, userUpdated user.UserUpdateDTO) error
	UpdatePassword(id uint, passwordUser user.UserUpdatePasswordDTO) error
}

type UserService struct {
	Db *gorm.DB
}

func (o *UserService) CreateUser(user user.UserCreationDTO) error {
	checkPassword := CheckUsernameAndPassword{}
	checkExist := CheckIfExist{Db: o.Db}
	saveUser := SaveUser{Db: o.Db}

	checkPassword.Next = &checkExist
	checkExist.Next = &saveUser

	if result, err := checkPassword.Execute(user); !result {
		return err
	}

	return nil
}

func (o *UserService) DeleteUser(id uint) error {
	db := o.Db
	user := entity.UserModel{}
	user.ID = id

	if result := db.Unscoped().Delete(&user); result.RowsAffected < 1 {
		return result.Error
	}

	return nil
}

func (o *UserService) GetAllUsers() ([]user.UserFrontDTO, error) {
	var users []entity.UserModel
	db := o.Db
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	usersDTO := make([]user.UserFrontDTO, len(users))

	for id, data := range users {
		var connectionEvents []entity.UserConnectionModel

		err := db.Model(&data).Association("Connection").Find(&connectionEvents)
		if err != nil {
			return nil, err
		}

		data.Connection = connectionEvents
		usersDTO[id] = data.UserModelToUserFrontDTO()
	}

	return usersDTO, nil
}

func (o *UserService) GetUserDetails(id uint) (*user.UserDetailsDTO, error) {
	db := o.Db
	var user entity.UserModel
	var userDetails entity.UserDetailsModel

	if result := db.Find(&user, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}

	db.Model(&user).Association("UserDetails").Find(&userDetails)

	user.UserDetails = userDetails
	result := user.UserModelToUserDetailsDTO()

	return &result, nil
}

func (o *UserService) SetConnection(username, password string) error {
	db := o.Db
	var userModel entity.UserModel

	if res := db.First(&userModel, "username = ?", username); res.RowsAffected == 0 {
		return errors.New("User not found")
	}

	err := db.Model(&userModel).Association("Connection").Append(&entity.UserConnectionModel{
		ConnectionDate: time.Now(),
	})

	if err != nil {
		return errors.New("Cannot set connection date")
	}

	if !helper.CheckPassword(password, &helper.HashedPassword{
		Digest: userModel.HashedPassword,
		Salt:   userModel.Salt}) {

		return errors.New("Password not found")
	}
	return nil
}

func (o *UserService) UpdateUser(id uint, userUpdated user.UserUpdateDTO) error {
	db := o.Db
	var userModel entity.UserModel

	if res := db.First(&userModel, "id = ?", id); res.Error != nil {
		return res.Error
	} else if res.RowsAffected <= 0 {
		return errors.New("Cannot find user")
	}

	err := db.Model(&userModel).Association("UserDetails").Append(&entity.UserDetailsModel{
		FirstName: userUpdated.FirstName,
		LastName:  userUpdated.LastName,
	})

	return err
}

func (o *UserService) UpdatePassword(id uint, passwordUser user.UserUpdatePasswordDTO) error {
	checkOldPassword := updatepassword.CheckOldPassword{
		Db: *o.Db,
	}
	checkPassword := updatepassword.CheckPassword{}
	updatePassword := updatepassword.UpdatePassword{
		Db: *o.Db,
	}

	checkOldPassword.Next = &checkPassword
	checkPassword.Next = &updatePassword

	_, err := checkOldPassword.Execute(id, passwordUser)
	return err
}

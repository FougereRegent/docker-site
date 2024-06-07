package userchecker

import (
	"docker-site/dto/user"
	"errors"
)

type CheckUsernameAndPassword struct {
	baseComponent
}

func (o *CheckUsernameAndPassword) Execute(user user.UserCreationDTO) (bool, error) {
	username := user.Name
	password := user.Password
	confirmPassword := user.ConfirmPassword

	if err := o.CheckUserNameCharacteristics(username); err != nil {
		return false, err
	}
	if err := o.CheckPasswordCharacteristics(password, confirmPassword); err != nil {
		return false, err
	}

	if o.Next != nil {
		return o.Next.Execute(user)
	}

	return true, nil
}

func (o *CheckUsernameAndPassword) CheckPasswordCharacteristics(password, confirmPassword string) error {
	if len(password) < 8 {
		return errors.New("Password lenght have to upper than 8 charactere")
	}
	if password != confirmPassword {
		return errors.New("Password and confirm Password are different")
	}
	return nil
}

func (o *CheckUsernameAndPassword) CheckUserNameCharacteristics(username string) error {
	if username == "" {
		return errors.New("Username cannot be empty")
	}
	if len(username) < 4 {
		return errors.New("Username lenght have to upper than 4 charactere")
	}
	return nil
}

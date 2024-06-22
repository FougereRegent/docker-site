package updatepassword

import (
	"docker-site/dto/user"
	userchecker "docker-site/service/user_checker"
)

type CheckPassword struct {
	baseComponent
}

func (o *CheckPassword) Execute(id uint, user user.UserUpdatePasswordDTO) (bool, error) {
	passwordChecker := userchecker.CheckUsernameAndPassword{}

	if err := passwordChecker.CheckPasswordCharacteristics(user.NewPassword, user.ConfirmNewPassword); err != nil {
		return false, err
	}

	if o.Next != nil {
		return o.Next.Execute(id, user)
	}
	return true, nil
}

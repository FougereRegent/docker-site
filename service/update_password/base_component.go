package updatepassword

import "docker-site/dto/user"

type IUpdaterPassword interface {
	Execute(id uint, userUpdatePassword user.UserUpdatePasswordDTO) (bool, error)
}

type baseComponent struct {
	Next IUpdaterPassword
}

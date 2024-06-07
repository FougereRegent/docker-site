package userchecker

import (
	"docker-site/dto/user"
)

type ICheckUser interface {
	Execute(user user.UserCreationDTO) (bool, error)
}

type baseComponent struct {
	Next ICheckUser
}

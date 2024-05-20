package dto

type UserFrontDTO struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

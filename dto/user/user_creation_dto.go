package user

type UserCreationDTO struct {
	Name            string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm-password"`
}

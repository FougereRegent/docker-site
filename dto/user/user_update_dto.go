package user

type UserUpdateDTO struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

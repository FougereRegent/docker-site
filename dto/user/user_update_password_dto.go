package user

type UserUpdatePasswordDTO struct {
	OldPassword        string `form:"old-password"`
	NewPassword        string `form:"new-password"`
	ConfirmNewPassword string `form:"confirm-password"`
}

package dto

type UserFrontDTO struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type TabDTO struct {
	Headers []string
	Values  [][]interface{}
}

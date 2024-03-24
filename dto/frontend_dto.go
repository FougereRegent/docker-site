package dto

type UserFrontDTO struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type TabDTO struct {
	UrlToScan string
	Headers   []string
	Values    [][]interface{}
}

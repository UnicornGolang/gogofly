package dto

// https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation
type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,email"`
	Password string `json:"password" binding:"required,capitalized"`
}

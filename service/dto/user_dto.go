package dto

// 各种不同的校验规则与校验信息自定义
// https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation
type UserLoginDTO struct {
	// 这里我们对客户端入参进行绑定的时候，对参数添加了校验，以及自定义参数校验信息
	// json: 表示绑定 json 中的字段名称
	// binding: 表示绑定的校验规则(可以使用官方提供的，也可以使用自定义的)
	// message: 表示默认的错误提示
	// required_err: 自定义信息
	Name     string `json:"name" binding:"required,email" message:"用户名填写错误" required_err:"用户名不能为空" email_err:"用户名必须符合邮箱规范"`
	Password string `json:"password" binding:"required,capitalized" message:"密码不能为空" capitalized_err:"密码必须符合title规则"`
}

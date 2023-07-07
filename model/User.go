package model

import (
	"gogofly/utils"

	"gorm.io/gorm"
)

type LoginUser struct {
	Id   uint
	Name string
}

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:64;not null"`
	RealName string `json:"real_name" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Mobile   string `json:"mobile" gorm:"size:11"`
	Email    string `json:"email" gorm:"size:128"`
	Password string `json:"-" gorm:"size:128;not null"`
}

// 要实现账户密码加密, 可以在模型类上设置对应的钩子函数
// 每次执行模型的查询，插入的时候会自动执行钩子函数，
// 从而实现账号密码的自动加密解密
// 参考文档 https://gorm.io/zh_CN/docs/hooks.html
// 以下是四个钩子函数在操作数据库是执行的时机,
// > 0.开始事务
// > 1.BeforeSave
// > 2.BeforeCreate
// > 3.关联前的 save
// > 4.插入记录至 db
// > 5.关联后的 save
// > 6.AfterCreate
// > 7.AfterSave
// > 8.提交或回滚事务
// --------------------------------------------------------------
// 要实现账号密码的加密，解密，
func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	m.Encrypt()
	return nil
}

func (m *User) Encrypt() {
	m.Password = utils.Encrypt(m.Password)
}

package service

import (
	"errors"
	"gogofly/dao"
	"gogofly/model"
	"gogofly/service/dto"
)

type UserService struct {
	BaseService
  Dao *dao.UserDao
}

var userService *UserService

func NewUserService() *UserService {
  if userService == nil {
    userService = &UserService{
      Dao: dao.NewUserDao(),
    }
  }
  return userService
}

// 
// 根据用户名查询用户, 校验密码是否正确, 实现登录功能
// 
func (m *UserService) Login(userDTO dto.UserLoginDTO) (model.User, error) {
  var errResult error
  user := m.Dao.GetUserByName(userDTO.Name)
  if user.ID == 0 {
    errResult = errors.New("invalid UserName")
  }
  // TODO: 对用户查到的密码加密后与数据库密码比较
  return user, errResult
}



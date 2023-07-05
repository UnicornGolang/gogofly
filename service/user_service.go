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

// 根据用户名查询用户, 校验密码是否正确, 实现登录功能
func (m *UserService) Login(userDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error
	user := m.Dao.GetUserByName(userDTO.Name)
	if user.ID == 0 {
		errResult = errors.New("invalid UserName")
	}
	// TODO: 对用户查到的密码加密后与数据库密码比较
	return user, errResult
}

// 添加用户的逻辑
func (m *UserService) AddUser(userDTO *dto.UserAddDTO) error {
	// 判断用户是否已经存在
	if m.Dao.CheckUserExist(userDTO.Name) {
		return errors.New("username Exist")
	}
	return m.Dao.AddUser(userDTO)
}

// 根据 ID 查找用户
func (m *UserService) GerUserById(commonDTO *dto.CommonDTO) (model.User, error) {
	return m.Dao.GetUserById(commonDTO.ID)
}

// 分页查找用户列表
func (m *UserService) GetUserList(userListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(userListDTO)
}

// 更新用户
func (m *UserService) UpdateUser(userUpdateDTO dto.UserUpdateDTO) error {
	if userUpdateDTO.ID == 0 {
		return errors.New("invalid user id")
	}

	return m.Dao.UpdateUser(&userUpdateDTO)
}

// 删除用户
func (m *UserService) DelUserById(commonDTO *dto.CommonDTO) error {
  return m.Dao.DelUserById(commonDTO.ID)
}


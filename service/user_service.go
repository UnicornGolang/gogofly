package service

import (
	"errors"
	"fmt"
	"gogofly/dao"
	global "gogofly/global/constants"
	"gogofly/model"
	"gogofly/service/dto"
	"gogofly/utils"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
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
func (m *UserService) Login(userDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
  var token string
	user, err := m.Dao.GetUserByName(userDTO.Name)
	if err != nil || user.ID == 0 {
		errResult = errors.New("invalid UserName")
    return user, token, errResult
  }
  // 校验密码
  if !utils.CompareHashAndPassword(user.Password, userDTO.Password) {
    errResult = errors.New("username or password is not correct") 
    return user, token, errResult
   }
  token, err = GenerateAndCacheLoginUserToken(user.ID, user.Name)
  if err != nil {
    errResult = fmt.Errorf("generate Token error: %s", err.Error())
  }
	return user, token, errResult
}

func GenerateAndCacheLoginUserToken(uid uint, name string) (string, error) {
  token, err := utils.GenerateToken(uid, name)
  if err == nil {
    err = global.RDB.Set(
      strings.ReplaceAll(global.LOGIN_USRE_TOKEN_REDIS_PREFIX, "{ID}", strconv.Itoa(int(uid))),
      token,
      time.Duration(viper.GetInt("jwt.tokenExpire"))*time.Minute,
    )
  }
  return token, err
}

// 添加用户的逻辑
func (m *UserService) AddUser(userDTO *dto.UserAddDTO) error {
	// 判断用户是否已经存在
	if m.Dao.CheckUserExist(userDTO.Name) {
		return errors.New("username Exist")
	}

	// 安装对应的加密包，系统的中的账号的密码不能明文保存，
	// 需要加密后保存到数据库中
	// go get -u golang.org/x/crypto/bcrypt
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

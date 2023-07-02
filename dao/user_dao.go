package dao

import (
	"gogofly/model"
)

type UserDao struct {
	BaseDao
}

var userDao *UserDao

func NewUserDao() *UserDao {
  if userDao == nil {
    userDao = &UserDao{
       NewBaseDao(),
    }
  }
  return userDao
}

func (m *UserDao) GetUserByName(name string) model.User {
  var user model.User 
  m.Orm.Model(&user).Where("name=?", name).Find(&user)
  return  user
}

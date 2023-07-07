package dao

import (
	"gogofly/model"
	"gogofly/service/dto"
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

func (m *UserDao) GetUserByName(name string) (model.User, error) {
	var user model.User
	err := m.Orm.Model(&user).Where("name=?", name).Find(&user).Error
	return user, err
}

func (m *UserDao) AddUser(userAddDTO *dto.UserAddDTO) error {
	var user model.User
	userAddDTO.ConvertToModel(&user)
	err := m.Orm.Save(&user).Error
	if err == nil {
		userAddDTO.ID = user.ID
		userAddDTO.Password = ""
	}
	return err
}

func (m *UserDao) CheckUserExist(username string) bool {
	var total int64
	m.Orm.Model(&model.User{}).Where("name=?", username).Count(&total)
	return total > 0
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	// 根据 ID 查找用户
	var user model.User
	err := m.Orm.First(&user, id).Error
	return user, err
}

func (m *UserDao) GetUserList(userListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var userList []model.User
	var total int64
	err := m.Orm.Model(&model.User{}).
		// 传入分页的回调函数
		Scopes(Paginate(userListDTO.Paginate)).
		Find(&userList).
		// 取消分页参数
		Offset(-1).Limit(-1).
		// 查找总数用以分页
		Count(&total).Error
	return userList, total, err
}

func (m *UserDao) UpdateUser(userUpdateDTO *dto.UserUpdateDTO) error {
	var user model.User
	m.Orm.First(&user, userUpdateDTO.ID)
	userUpdateDTO.ConvertToModel(&user)

	// Save 方法直接覆盖所有字段，不管传入的对象是否存在零值
	// ----------------------------------------------------------------------------------------
	// u := User
	// u.Name:"alex"
	// db.Model(&user).Save(&u)
	//
	// Update 更新单个字段
	// ----------------------------------------------------------------------------------------
	// 1. db.Model(&user)Update("name", "hello")
	// 2. db.Model(&user).Where("active = ?").Update("name", "hello")
	//
	// Updates 更新多个字段 Updates("map")
	// ----------------------------------------------------------------------------------------
	// 1. db.Model(&user).Updates(User{Name:"hello", Age: 18, Active: false})
	// 2. db.Modul(&user).Updates(map[string]any{"name": "hello", "age":18, "active": false})
	return m.Orm.Save(&user).Error
}

// 删除用户
func (m *UserDao) DelUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}

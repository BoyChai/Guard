package dao

import (
	"errors"

	"github.com/BoyChai/Guard/utils"
)

// CreateUser 创建用户
func (d *dao) CreateUser(name string, pass string) (uint, error) {
	user := User{Name: name, Pass: utils.CalculateMD5Hash(pass)}
	err := Dao.db.Create(&user).Error
	return user.ID, err
}

// CheckUser 用户校验
func (d *dao) CheckUser(name string, pass string) (uint, error) {
	var user User
	err := Dao.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return 0, err
	}
	if user.Pass != utils.CalculateMD5Hash(pass) {
		return 0, err
	}
	return user.ID, nil
}

// DeleteUserByID 删除用户
func (d *dao) DeleteUserByID(id uint) error {
	if id == 1 { // 不允许删除默认的超级用户
		return errors.New("不允许删除默认的超级用户")
	}
	return Dao.db.Delete(&User{}, id).Error
}

// 列出用户
func (d *dao) ListUser() ([]User, error) {
	var users []User
	err := Dao.db.Find(&users).Error
	return users, err
}

// 修改密码
func (d *dao) ChangePass(id uint, pass string) error {
	return Dao.db.Model(&User{}).Where("id = ?", id).Update("pass", utils.CalculateMD5Hash(pass)).Error
}

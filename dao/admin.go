package dao

import (
	"auth_frame/dto"
	"auth_frame/public"
	"errors"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Uid      int    `json:"uid" gorm:"unique" description:"唯一id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func (admin *Admin) Find(db *gorm.DB, uid int) (*Admin, error) {
	result := db.Where("uid = ?", uid).First(admin)
	return admin, result.Error
}

func (admin *Admin) Update(db *gorm.DB) error {
	oldAdmin := &Admin{}
	row := db.Where("uid = ?", admin.Uid).First(oldAdmin)
	if row.Error != nil {
		return row.Error
	}
	result := db.Model(oldAdmin).Updates(admin)
	return result.Error

}

func (admin *Admin) Delete(db *gorm.DB, uid int) error {
	result := db.Where("uid = ?", uid).Delete(admin)
	return result.Error
}

func UidCheck(db *gorm.DB, uid int) bool {
	admin := &Admin{}
	db.Where("uid = ?", uid).First(admin)
	if admin.ID != 0 {
		return false
	}
	return true
}
func (admin *Admin) LoginCheck(db *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	db.Where("username = ?", param.Username).First(admin)

	if admin.ID == 0 {
		return nil, errors.New("用户名不存在")
	}
	if !public.ValidPassword(admin.Password, param.Password) {
		return nil, errors.New("密码错误，请重新输入")
	}
	return admin, nil
}

func RegisterCheck(db *gorm.DB, param *dto.RegisterInput) error {
	admin := &Admin{}
	db.Where("username = ?", param.Username).First(admin)
	if admin.ID != 0 {
		return errors.New("用户已存在")
	}
	return nil
}

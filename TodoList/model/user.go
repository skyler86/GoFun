package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`		//将用户名设置成唯一
	PasswordDigest string	// 数据库存储的密码是密文，也就是加密后的密码
}

// 实现密码加密
func (user *User) SetPassword(password string) error {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),12)		//cost为加密难度
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 实现密码验证
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest),[]byte(password))
	return err==nil
}
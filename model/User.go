package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

//// BeforeCreate 创建前的钩子
//func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	u.Role = 2
//	return nil
//}
//
//// BeforeUpdate 更新前的钩子
//func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}

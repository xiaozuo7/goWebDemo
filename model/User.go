package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	BaseModel `json:"-"`
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

//BeforeCreate 创建前的钩子
//func (u *User) BeforeCreate(_ *gorm.DB) {
//	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
//
//}

//BeforeUpdate 更新前的钩子
//func (u *User) BeforeUpdate(_ *gorm.DB) {
//	u.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
//
//}


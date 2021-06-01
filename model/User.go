package model

import (
	"goWebDemo/utils/errmsg"
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

func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUserNameExists
	}
	return errmsg.Success
}

func CreateUser(data *User) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

func CheckLogin(username string, password string) (User, int) {
	var user User

	db.Where("username = ?", username).First(&user)

	PasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ErrorUserNotExists
	}

	if PasswordErr != nil {
		return user, errmsg.ErrorPasswordInvalid
	}

	if user.Role != 1 {
		return user, errmsg.ErrorPermissionDenied
	}
	return user, errmsg.Success
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

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

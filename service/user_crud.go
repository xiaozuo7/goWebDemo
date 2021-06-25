package service

import (
	"goWebDemo/model"
	"goWebDemo/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
)

var err error

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user model.User
	model.Db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUserNameExists
	}
	return errmsg.Success
}

// CreateUser 创建用户
func CreateUser(data *model.User) (code int) {
	data.Password = model.ScryptPw(data.Password)  // 密码加密
	err = model.Db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// CheckLogin 后台登录检验
func CheckLogin(username string, password string) (model.User, int) {
	var user model.User

	model.Db.Where("username = ?", username).First(&user)

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

//CheckLoginFront 前台用户校验
func CheckLoginFront(username string, password string) (model.User, int) {
	var user model.User

	model.Db.Where("username = ?", username).First(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ErrorUserNotExists
	}
	if err != nil {
		return user, errmsg.ErrorPasswordInvalid
	}
	return user, errmsg.Success
}

// EditUser 编辑用户
func EditUser(id int, data *model.User) int {
	var user model.User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = model.Db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// CheckUpUser 更新查询
func CheckUpUser(id int, username string) (code int) {
	var user model.User
	model.Db.Select("id, username").Where("username = ?", username).First(&user)
	if user.ID == uint(id) {
		return errmsg.Success
	}
	if user.ID > 0 {
		return errmsg.ErrorUserNameExists
	}
	return errmsg.Success
}

// GetUser 查询单个用户
func GetUser(id int) (model.User, int) {
	var user model.User
	err = model.Db.Select("created_at, updated_at, id, username, role").Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ErrorUserNotExists
	}
	return user, errmsg.Success
}

// GetUserList 获取用户列表
func GetUserList(username string, pageSize int, pageNum int) ([]model.User, int64) {
	var users []model.User
	var total int64

	if username != "" {
		model.Db.Select("id, username, role").Where(
			"username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		model.Db.Model(&users).Where("username LIKE ?", username+"%").Count(&total)
		return users, total
	}
	model.Db.Select("id, username, role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	model.Db.Model(&users).Count(&total)

	return users, total
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user model.User
	err = model.Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// ChangePassword 修改密码
func ChangePassword(id int, password string) int {
	var user model.User
	hashPassword := model.ScryptPw(password)
	//var maps = make(map[string]interface{})
	//maps["password"] = hashPassword
	err = model.Db.Model(&user).Select("password").Where("id = ?", id).Update("password", hashPassword).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

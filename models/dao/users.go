package dao

import (
	"IRIS_WEB/models/dto"
	"IRIS_WEB/utility/db"
	"time"
)

// 创建用户
func CreateUser(user *dto.UserDTO) (err error) {
	res := db.GetMysql().Create(&user)
	err = res.Error
	return
}

// 更新用户
func UpdateUser(user *dto.UserDTO) (err error) {
	res := db.GetMysql().Model(&user).Updates(dto.UserDTO{
		UserName:     user.UserName,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		Status:       user.Status,
		UpdatedAt:    time.Now(),
	})
	err = res.Error
	return
}

// 删除用户
func DeleteUser(userId int) (err error) {
	user := new(dto.UserDTO)
	res := db.GetMysql().Delete(&user, userId)
	err = res.Error
	return
}

// 根据用户ID查询
func QueryUserById(userId int) (user *dto.UserDTO, err error) {
	user = new(dto.UserDTO) // 如果不通过new关键字实例化就会报错: unsupported destination, should be slice or struct
	res := db.GetMysql().First(&user, userId)
	err = res.Error
	return
}

// 根据用户名模糊查询
//func QueryUsersByName(userName string) (users []*dto.UserDTO, err error) {
//	res := db.GetMysql().Where("user_name like ?", "%" + userName + "%").Find(&users)
//	err = res.Error
//	return
//}

// 查询所有用户
func QueryAllUsers() (users []*dto.UserDTO, err error) {
	res := db.GetMysql().Find(&users)
	err = res.Error
	return
}

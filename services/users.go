package services

import (
	"IRIS_WEB/errors"
	"IRIS_WEB/models/dao"
	"IRIS_WEB/models/dto"
)

func CreateUser(user *dto.UserDTO) error {
	var err error
	if err = dao.CreateUser(user); err != nil {
		return errors.DBError("create user failed", err)
	}

	return nil
}

func UpdateUser(user *dto.UserDTO) error {
	var err error
	if err = dao.UpdateUser(user); err != nil {
		return errors.DBError("update user failed", err)
	}

	return nil
}

func DeleteUser(userId int) error {
	// 根据ID删除用户
	var err error
	if err = dao.DeleteUser(userId); err != nil {
		return errors.DBError("delete user by id", err)
	}

	return nil
}

func FetchUserById(userId int) (*dto.UserDTO, error) {
	// 根据ID查询用户
	var err error
	var user *dto.UserDTO
	if user, err = dao.QueryUserById(userId); err != nil {
		return nil, errors.DBError("query user by id", err)
	}

	return user, nil
}

func FetchAllUsers() ([]*dto.UserDTO, error) {
	// 查询所有用户
	var err error
	var users []*dto.UserDTO
	if users, err = dao.QueryAllUsers(); err != nil {
		return nil, errors.DBError("query all users", err)
	}

	return users, nil
}

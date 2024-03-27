package service

import (
	"main/db"
	"main/entity"
)

func CreateUser(user *entity.User) (err error) {
	if err = db.GetDb().Create(user).Error; err != nil {
		return err
	}
	return
}

func GetAllUser() (userList []*entity.User, err error) {
	if err := db.GetDb().Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteUserById(id string) (err error) {
	err = db.GetDb().Where("id=?", id).Delete(&entity.User{}).Error
	return
}

func GetUserById(id string) (user *entity.User, err error) {
	if err = db.GetDb().Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *entity.User) (err error) {
	err = db.GetDb().Save(user).Error
	return
}

package db

import (
	"errors"
	"offer_tiktok/pkg/constants"
)

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func QueryUser(userName string) (*User, error) {
	var user User
	if err := DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserID(userID int64) (ok bool, err error) {
	var user User
	if err := DB.Where("ID = ?", userID).Find(&user).Error; err != nil {
		return false, err
	}
	if userID == 0 {
		err := errors.New("userID not found")
		return false, err
	}
	return true, nil
}

func VerifyUser(userName string, password string) (int64, error) {
	var user User
	if err := DB.Where("user_name = ? AND password = ?", userName, password).Find(&user).Error; err != nil {
		return 0, err
	}
	if user.ID == 0 {
		err := errors.New("username or password not verified")
		return user.ID, err
	}
	return user.ID, nil
}

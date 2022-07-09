package service

import (
	"blog-go/pkg/db"
	"blog-go/pkg/model"
)

func CreateUser(user *model.UserInfo) error {
	user.Status=1
	err:=db.CreateUser(user)
	return err
}
func UpdateUser(user *model.UserInfo) error  {
	err:=db.UpdateUser(user)
	return err
}
func DeleteUser(user *model.UserInfo) error  {
	user.Status=0
	err:=db.DeleteUser(user)
	return err
}
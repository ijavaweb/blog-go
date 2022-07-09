package db

import "blog-go/pkg/model"

func CreateUser(user *model.UserInfo)  error{
	db:=DB.Table("user_info").Create(user)
	return db.Error
}

func UpdateUser(user *model.UserInfo)  error{
	db:=DB.Table("user_info").Update(user).Where("id=?",user.Id)
	return db.Error
}

func DeleteUser(user *model.UserInfo)  error{
	db:=DB.Table("user_info").Updates(user).Where("id=?",user.Id)
	return db.Error
}

package db

import (
	"blog-go/pkg/model"
)

func CreateCategory(c *model.Category)  error{
	db:=DB.Table("category").Create(c)
	return db.Error
}

func UpdateCategory(c *model.Category)  error{
	db:=DB.Table("category").Update(c).Where("id=?",c.Id)
	return db.Error
}

func DeleteCategory(c *model.Category)  error{
	db:=DB.Table("category").Updates(c).Where("id=?",c.Id)
	return db.Error
}

func GetCategoryList() ([]model.Result,error)  {
	result:=make([]model.Result,0)
	db:=DB.Table("article").Group("category").Select("category").Find(&result)
	if db.Error!= nil {
		return nil,db.Error
	}
	return result,nil
}
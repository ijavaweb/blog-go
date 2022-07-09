package service

import (
	"blog-go/pkg/db"
	"blog-go/pkg/model"
)

func CreateCategory(c *model.Category) error {
	c.Status=1
	err:=db.CreateCategory(c)
	return err
}
func UpdateCategory(c *model.Category) error  {
	err:=db.UpdateCategory(c)
	return err
}
func DeleteCategory(c *model.Category) error  {
	c.Status=0
	err:=db.DeleteCategory(c)
	return err
}
func GetCategoryList() ([]string,error)  {
	result,err:=db.GetCategoryList()
	if err != nil {
		return nil,err
	}
	category:=make([]string,0)
	for _, r := range result {
		if r.Category != "" {
			category = append(category,r.Category)
		}
	}
	return category,nil
}
package service

import "blog-go/pkg/db"

func GetVisitorById(id int32) (int32,error)  {
	return db.GetVisitorById(id)
}
func GetVisitorAmount() (int32,error) {
	return db.GetVisitorAmount()
}
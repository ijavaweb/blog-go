package db

func GetVisitorById(id int32) (int32,error){
	count:=0
	db:=DB.Table("visitor").Count(&count).Where("id=?",id)
	return int32(count),db.Error
}

func GetVisitorAmount() (int32,error){
	count:=0
	db:=DB.Table("visitor").Count(&count)
	return int32(count),db.Error
}
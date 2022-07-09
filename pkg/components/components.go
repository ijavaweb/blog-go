package components

import (
	"blog-go/pkg/db"
)

func InitComponents() {
	db.InitMySQL()
}

package connection

import (
	"test-d-2/structs"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
	Err error
)

func Connect(){
	DB, Err = gorm.Open("mysql", "fajar:quNrJ1NLoDRuvRNs@/fajar?charset=utf8&parseTime=True")

	if Err != nil{
		log.Println("Connection failed", Err)
	}else{
		log.Println("Server upn and running")
	}

	DB.AutoMigrate(&structs.User{})
	DB.AutoMigrate(&structs.RiskProfile{})
}
package config


import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)
var DB *gorm.DB
func ConnectDB(){
	dsn := "database url."
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("database not connected bro")
		return
	}
	DB = db
	fmt.Println("database connected succesfully bro")


}
package config


import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)
var DB *gorm.DB
func ConnectDB(){
	
	dsn := "host=localhost user=postgres password=yourpassword dbname=gormdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("database not connected bro")
		return
	}
	DB = db
	fmt.Println("database connected succesfully bro")


}
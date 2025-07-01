package main 

import (
	"gorm.io/grom"
	"fmt"
	"task_scheduler/models"
	"task_scheduler/config"
)


func main(){
	//connect to db
	config.ConnectDB()
	//auto migrate the db
	config.DB.AutoMigrate(&models.Task{})

	
}
package commands

import (
	"fmt"
	"task_scheduler/config"
	"task_scheduler/models"
)

func Create(Message string, delay int){
	task := models.Task {
		Message:Message,
		Delay:delay,
		Executed: false,
	}
	//i have succesfully created a struct.now neeed to store in db bolte.
	res := config.DB.Create(&task)
	if res.Error != nil{
		fmt.Print("honey singh not created in db beybye")
		return
	}
	fmt.Print("honey singh saved in db succesfully.")


}
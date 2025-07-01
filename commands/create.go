package commands

import (
	"fmt"
	"task_scheduler/config"
	"task_scheduler/models"
)

func Create(message string, delay int){
	task := models.Task {
		Message:message,
		Delay:delay,
		Executed: false,
	}
	//i have succesfully created a struct.now neeed to store in db bolte.
	res := config.DB.Create(&task)
	if res.Error != nil{
		fmt.Print("honey singh not created in db beybye")
		return
	}
	fmt.Printf("✅ Task created: \"%s\" will run after %d seconds (ID: %d)\n", message, delay, task.ID)
}
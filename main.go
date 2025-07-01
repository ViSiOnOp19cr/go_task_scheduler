package main 

import (
	
	"task_scheduler/models"
	"task_scheduler/config"
	"task_scheduler/commands"
	"os"
	"strconv"
	"fmt"
)


func main(){
	//connect to db
	config.ConnectDB()
	//auto migrate the db
	config.DB.AutoMigrate(&models.Task{})

	// now i have to read the command from terminal for that we need to use os

	args := os.Args

	if len(args) < 2 {
		fmt.Print("nah nah missing args try again later")
		return
	}
	switch args[1]{
	case "create":
		if len(args) < 4 {
			fmt.Println(" Usage: create <message> <delay>")
			return
		}
		//call create command func
		message :=args[2]
		delay, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("❌ Delay must be an integer (in seconds)")
			return
		}
		commands.Create(message,delay)
	default:
		fmt.Println("❌ Unknown command")
	}

}
package main 

import (
	"log"
	"flag"
	"strings"
	"slices"
)

type Task struct {
    ID      int
    Title   string
    Done    bool
}

var TASKS []Task 

func main() {
	var command string
	var description string
	var duedate string
	var id int
	flag.StringVar(&command,"cmd","","Provide the action you want to perform")
	flag.StringVar(&description,"desp","Shubham","Describe your task")
	flag.StringVar(&duedate,"ddate","Shubham","duedate of the task")
	flag.IntVar(&id, "id",0,"The id of the string you want to mark done")
	
	flag.Parse()
	


	tasks_name := []string{"add","list","done"}

	command = strings.TrimSpace(
		strings.ToLower(command))
	
	if ! slices.Contains(tasks_name, command){
		log.Println("can not find command to execute")
		log.Println("Please provide one of the following commands",tasks_name)
		flag.Usage()
	}

	if command == "add"{
		log.Println("You have asked to add a task")
			newtask := Task{
				ID:  len(TASKS) + 1,
				Title: description,
				Done: false,
			}
			TASKS = append(TASKS,newtask)
	} else if command == "list" {
		log.Println("You have asked to show the list of tasks")
		for _, t := range TASKS {
			log.Printf("ID: %d | Title: %s | Done: %t\n", t.ID, t.Title, t.Done)
		}
	} else if command == "done" {
		log.Println("You have asked to mark a task as complete")
		is_done := markTaskDone(id)
		if ! is_done{
			log.Println("Failed !! Please provide a valid id")
		}else{
			log.Println("Task with id ",id," marked done")
		}
	}

}

func markTaskDone(id int) bool {
	for i, task := range TASKS {
		if task.ID == id {
			TASKS[i].Done = true
			return true // task marked as done
		}
	}
	return false // task not found
}
package main

import "fmt"

// func tasks(task){
	
// }

func main() {
	var taskOne = "Watch go lang tutorial"
	var taskItems = []string{
		taskOne,
		taskOne,
		taskOne,
	}

	fmt.Println("Welcome to TODOist!")
	
	fmt.Println("Todo List:")
	fmt.Println(taskItems[2])
}
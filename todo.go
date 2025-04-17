package main

import "fmt"

// Functions in Go

func printTasks(taskItemsSlice []string){ // printing tasks 
	fmt.Println("List of my Todos")
	for index, task := range taskItemsSlice {
		index = index + 1
		fmt.Printf("%d. %s\n", index, task)
	}
}

//When returning a value, make sure to specify what type of value are you returning, in the case for the functions below, they are returning a string slice 
func addTask(task string, taskItemsSlice []string) []string { // function to add a task to the array or slice
	var updatedItems = append(taskItemsSlice, task) // append the task to the slice and return the new slice
	printTasks(updatedItems)
	return updatedItems
} // end of addTask function

func removeTask(index int, taskItemsSlice []string) []string { // function to remove a task given the index 
	index = index - 1
	var updatedItems = append(taskItemsSlice[:index], taskItemsSlice[index+1:]...) // removing a value from a specified index
	printTasks(updatedItems)
	return updatedItems
}

func main() {
	// Printing statements 
	fmt.Println("Welcome to TODOist!")

	// Declaring variables 
	var taskOne = "Watch go lang tutorial in a slice" // declaring a variable (you should be able to lock in what the data type for a variable can be as well)
	var taskTwo = "Know what the difference betweeen an array and a slice is" // declaring second variable 
	var taskThree = "Know how to declare different kinds of arrays and slices"
	var taskFour = "This is task 4"
	// Slices and Arrays 
	var taskItemsSlice = []string{
		taskOne,
		taskTwo,
		taskThree,
	} // declaring a string slice, can have unlimited values inside it
	
	fmt.Println("Todo List:")
	printTasks(taskItemsSlice)
	fmt.Println("Added a new task")
	taskItemsSlice = addTask(taskFour, taskItemsSlice)
	

	fmt.Println("Removed a task")
	taskItemsSlice = removeTask(3, taskItemsSlice)
} // end of main function

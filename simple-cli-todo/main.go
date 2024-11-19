package main

import ("fmt";"bufio" ;"os")

type Task struct{
	id int
	task string
	isCompleted bool
}

func listOperations(){
	fmt.Println("Enter Operation")
	fmt.Println("1. Add Task")
	fmt.Println("2. Delete Task")
	fmt.Println("3. List Task")
	fmt.Println("4. Exit")
}

func addTask(taskList *[]Task){
	reader := bufio.NewReader(os.Stdin)	
	fmt.Println("Enter Task")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] 	
	*taskList = append(*taskList, Task{id:len(*taskList)+1, task: input, isCompleted:false} )
	fmt.Println("Task Added")
}

func deleteTask(taskList *[]Task){
	listTasks(taskList)
	fmt.Println("Enter the id of the task to delete")	
	var taskID int
	_,err := fmt.Scanln(&taskID)
	if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
		}	
	var temp []Task
	for _, task := range *taskList{
		if task.id != taskID{
			temp = append(temp,task)
		}
	}
	*taskList = temp
}

func listTasks(taskList *[]Task){
	fmt.Println("The Tasks are:")
	for _, task :=range *taskList{
		fmt.Println("%v",task)
	}
}

func exit(){
	os.Exit(0)
} 


func main(){
		var taskList []Task
var operation int
	for{
	fmt.Println("Simple ToDo Cli App")
	listOperations()
	_, err := fmt.Scanln(&operation) // Use Scanln to read the operation
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
	switch operation{
	case 1:
		addTask(&taskList)
	case 2:
		deleteTask(&taskList)
	case 3:
		listTasks(&taskList)
	case 4:
		exit()
	default:
		fmt.Println("Invalid Operation")
	}
}
}
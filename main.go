package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"toDoApp/todolist"
)

func main() {
	todoList := todolist.NewToDoList()
	scanner := bufio.NewScanner(os.Stdin)

	todoList.AutoLoad()

	for {
		fmt.Println("\nTo-Do App")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Export Tasks")
		fmt.Println("6. Import Tasks")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: \n")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter due date (YYYY-MM-DD): ")
			scanner.Scan()
			dueDateInput := scanner.Text()
			dueDate, err := time.Parse("2006-01-02", dueDateInput)
			if err != nil {
				fmt.Println("Invalid due date format. Please use YYYY-MM-DD.")
				continue
			}

			todoList.AddTask(title, dueDate)
			todoList.AutoSave()

		case "2":
			todoList.ViewTasks()
		case "3":
			fmt.Print("Enter task ID to complete: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task ID.")
				continue
			}
			if err := todoList.CompleteTask(id); err != nil {
				fmt.Println(err)
			}
			todoList.AutoSave()

		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task ID.")
				continue
			}
			if err := todoList.DeleteTask(id); err != nil {
				fmt.Println(err)
			}
			todoList.AutoSave()

		case "5":
			fmt.Print("Enter filename to export tasks: ")
			scanner.Scan()
			filename := scanner.Text()
			if err := todoList.SaveTasks(filename); err != nil {
				fmt.Println(err)
			}

		case "6":
			fmt.Print("Enter filename to import tasks: ")
			scanner.Scan()
			filename := scanner.Text()
			if err := todoList.LoadTasks(filename); err != nil {
				fmt.Println(err)
			}
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

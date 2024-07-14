package todolist

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ToDoList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func NewToDoList() *ToDoList {
	return &ToDoList{NextID: 1}
}

func (todolist *ToDoList) AddTask(title string, dueDate time.Time) {
	task := Task{
		ID:        todolist.NextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		DueDate:   dueDate,
	}
	todolist.Tasks = append(todolist.Tasks, task)
	todolist.NextID++
	fmt.Println("Task added successfully!")
}

func (todolist *ToDoList) ViewTasks() {
	if len(todolist.Tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	for _, task := range todolist.Tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s] (Created: %s, Due: %s)\n", task.ID, task.Title, status, task.CreatedAt.Format(time.RFC1123), task.DueDate.Format(time.RFC1123))
	}
}
func (todolist *ToDoList) CompleteTask(id int) error {
	for i, task := range todolist.Tasks {
		if task.ID == id {
			todolist.Tasks[i].Done = true
			fmt.Println("Task marked as completed!")
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (todolist *ToDoList) DeleteTask(id int) error {
	for i, task := range todolist.Tasks {
		if task.ID == id {
			todolist.Tasks = append(todolist.Tasks[:i], todolist.Tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (todolist *ToDoList) SaveTasks(filename string) error {
	data, err := json.Marshal(todolist)
	if err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

func (todolist *ToDoList) AutoSave() error {
	if err := todolist.SaveTasks("auto"); err != nil {
		return fmt.Errorf("error autosaving to file auto")
	}
	return nil
}

func (todolist *ToDoList) LoadTasks(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	err = json.Unmarshal(data, todolist)
	if err != nil {
		return fmt.Errorf("error loading tasks: %w", err)
	}

	// Update NextID to avoid duplicate IDs after loading
	maxID := 0
	for _, task := range todolist.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	todolist.NextID = maxID + 1
	return nil
}

func (todolist *ToDoList) AutoLoad() error {
	if _, err := os.Stat("auto"); err != nil {
		os.Create("auto")
	}
	if err := todolist.LoadTasks("auto"); err != nil {
		return fmt.Errorf("error autoloading form file auto")
	}
	return nil
}

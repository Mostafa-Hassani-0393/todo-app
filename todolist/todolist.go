package todolist

// This is where the Tasks are saved
type ToDoList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

// Constructor for ToDoList
func NewToDoList() *ToDoList {
	return &ToDoList{NextID: 1}
}

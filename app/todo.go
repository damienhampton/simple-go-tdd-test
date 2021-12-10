package app

type Todo struct {
	Description string
	Complete    bool
	Id          int
}

func NewTodo(description string, index int) Todo {
	return Todo{description, false, index}
}

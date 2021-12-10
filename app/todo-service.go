package app

type TodoService struct {
	db     DBInterface
	nextId int
}

func NewTodoService(db DBInterface) TodoService {
	return TodoService{db: db, nextId: 0}
}

func (t *TodoService) Add(description string) Todo {
	newTodo := NewTodo(description, t.nextId)
	t.nextId = t.nextId + 1
	t.db.Insert(newTodo)
	return newTodo
}

func (t TodoService) List() []Todo {
	return t.db.List()
}

func (t TodoService) Complete(id int) Todo {
	todo := t.db.GetById(id)
	todo.Complete = true
	t.db.Update(todo)
	return todo
}

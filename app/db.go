package app

type InMemDB struct {
	todos []Todo
}

func (f *InMemDB) Insert(todo Todo) {
	f.todos = append(f.todos, todo)
}
func (f *InMemDB) Update(todo Todo) {
	f.todos[todo.Id] = todo
}
func (f InMemDB) GetById(id int) Todo {
	return f.todos[id]
}
func (f InMemDB) List() []Todo {
	return f.todos
}

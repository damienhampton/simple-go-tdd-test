package app

type DBInterface interface {
	Insert(todo Todo)
	GetById(id int) Todo
	Update(todo Todo)
	List() []Todo
}

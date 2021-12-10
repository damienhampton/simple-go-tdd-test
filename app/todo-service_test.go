package app_test

import (
	"testing"

	"dh.com/test1/app"
)

func TestTodoService(t *testing.T) {
	t.Run("Add two todos and return it", func(t *testing.T) {
		fakeDb := &app.InMemDB{}
		service := app.NewTodoService(fakeDb)
		service.Add("Stroke cat")
		service.Add("Feed cat")
		todos := service.List()
		if todos[0].Description != "Stroke cat" {
			t.Errorf("Should be Stroke cat, got %v", todos[0])
		}
		if todos[0].Complete != false {
			t.Errorf("Task should not be complete")
		}
		if todos[1].Description != "Feed cat" {
			t.Errorf("Should be Feed cat, got %v", todos[1])
		}
		if todos[1].Complete != false {
			t.Errorf("Task should not be complete")
		}
	})
	t.Run("Complete a todo", func(t *testing.T) {
		fakeDb := &app.InMemDB{}
		service := app.NewTodoService(fakeDb)
		todo := service.Add("Stroke cat")
		if todo.Complete != false {
			t.Errorf("Task should not be complete")
		}
		updatedTodo := service.Complete(todo.Id)
		if updatedTodo.Complete != true {
			t.Errorf("Task should be complete")
		}
	})
}

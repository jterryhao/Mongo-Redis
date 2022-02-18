package dao

import "github.com/jterryhao/Mongo-Redis/model"

type ToDoDataAccessor interface {
	CreateTodoItem(t *model.ToDoItem) error

	GetTodoItem(id string) (t *model.ToDoItem, err error)

	UpdateTodoItem(t *model.ToDoItem) error

	DeleteTodoItem(id string) error
}

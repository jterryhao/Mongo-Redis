package model

import "github.com/kamva/mgm/v3"

type ToDoItem struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Description      string `json:"description" bson:"description"`
}

func NewToDoItem(description string) *ToDoItem {
	return &ToDoItem{
		Description: description,
	}
}

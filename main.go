package main

import (
	"fmt"
	"github.com/jterryhao/Mongo-Redis/dao"
	"github.com/jterryhao/Mongo-Redis/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	//mongoClient, err := dao.NewMongoClient()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	redisClient := dao.NewRedisClient()

	var client dao.ToDoDataAccessor
	client = redisClient

	err := client.Ping()

	numOfTries := 2

	for i := 0; i < numOfTries; i++ {
		todo := model.NewToDoItem("Mongo and Redis")
		todo.SetID(primitive.NewObjectID())

		err = client.CreateTodoItem(todo)
		if err != nil {
			fmt.Printf("error creating %+v: %v\n", todo, err)
			return
		}
		fmt.Printf("created todo item with id: %v\n", todo.ID)

		readTodo, err := client.GetTodoItem(todo.ID.Hex())
		if err != nil {
			fmt.Printf("error getting todo item, %s, err: %v\n", todo.ID, err)
			return
		}
		fmt.Printf("read todo item: %+v\n", readTodo)

		todo.Description = "Redis and Mongo"
		err = client.UpdateTodoItem(todo)
		if err != nil {
			fmt.Printf("error updating %+v: %v\n", todo, err)
			return
		}
		fmt.Printf("updated todo item: %+v\n", todo)

		err = client.DeleteTodoItem(todo.ID.Hex())
		if err != nil {
			fmt.Printf("error delete todo item, %s, err: %v\n", todo.ID, err)
			return
		}
		fmt.Printf("deleted todo item: %s\n", todo.ID)

		readTodo, err = client.GetTodoItem(todo.ID.Hex())
		if readTodo != nil {
			fmt.Printf("got non-nil item, %s, err: %v\n", todo.ID, err)
			return
		}
		fmt.Printf("item deleted")
	}
}

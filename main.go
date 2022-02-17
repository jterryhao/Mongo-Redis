package main

import (
	"fmt"
	"github.com/jterryhao/Mongo-Redis/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	testConfig := &mgm.Config{CtxTimeout: time.Second * 30}
	connString := "mongodb+srv://terry-test:terry-test@terry-test-cluster.r5h6f.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	err := mgm.SetDefaultConfig(testConfig,
		"terry-test",
		options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Println(err)
	}

	todo := model.NewToDoItem("Mongo and Redis")

	coll := mgm.CollectionByName("todo-list")
	err = coll.Create(todo)
	if err != nil {
		fmt.Println(err)
	}
}

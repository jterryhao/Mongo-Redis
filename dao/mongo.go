package dao

import (
	"context"
	"github.com/jterryhao/Mongo-Redis/model"
	"github.com/jterryhao/Mongo-Redis/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoClient struct {
	Client *mongo.Client
}

func NewMongoClient() (*mongoClient, error) {
	defer utils.TimeTrack(time.Now(), "init mongo Client")
	testConfig := &mgm.Config{CtxTimeout: time.Second * 30}
	connString := "mongodb+srv://terry-test:terry-test@terry-test-cluster.r5h6f.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	err := mgm.SetDefaultConfig(testConfig,
		"terry-test",
		options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}

	_, mc, _, err := mgm.DefaultConfigs()

	return &mongoClient{
		Client: mc,
	}, nil
}

func (c *mongoClient) CreateTodoItem(t *model.ToDoItem) error {
	defer utils.TimeTrack(time.Now(), "create")

	coll := mgm.CollectionByName("todo-list")
	return coll.Create(t)
}

func (c *mongoClient) GetTodoItem(id string) (t *model.ToDoItem, err error) {
	defer utils.TimeTrack(time.Now(), "read")

	t = &model.ToDoItem{}
	coll := mgm.CollectionByName("todo-list")

	err = coll.FindByID(id, t)
	return
}

func (c *mongoClient) UpdateTodoItem(t *model.ToDoItem) error {
	defer utils.TimeTrack(time.Now(), "update")

	coll := mgm.CollectionByName("todo-list")

	return coll.Update(t)
}

func (c *mongoClient) DeleteTodoItem(id string) error {
	defer utils.TimeTrack(time.Now(), "delete")

	coll := mgm.CollectionByName("todo-list")

	t := &model.ToDoItem{}
	err := coll.FindByID(id, t)
	if err != nil {
		return err
	}

	return coll.Delete(t)
}

func (c *mongoClient) Ping() error {
	defer utils.TimeTrack(time.Now(), "ping")

	return c.Client.Ping(context.Background(), nil)
}

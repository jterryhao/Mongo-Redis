package dao

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/jterryhao/Mongo-Redis/model"
	"github.com/jterryhao/Mongo-Redis/utils"
	"time"
)

type redisClient struct {
	Client *redis.Client
}

func NewRedisClient() *redisClient {
	defer utils.TimeTrack(time.Now(), "init redis Client")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-12971.c98.us-east-1-4.ec2.cloud.redislabs.com:12971",
		Password: "mongo-redis", // no password set
		DB:       0,             // use default DB
	})
	return &redisClient{Client: rdb}
}

func (r *redisClient) CreateTodoItem(t *model.ToDoItem) error {
	defer utils.TimeTrack(time.Now(), "create in cache")
	entry, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return r.Client.Set(context.Background(), t.ID.Hex(), entry, 0).Err()
}

func (r *redisClient) GetTodoItem(id string) (t *model.ToDoItem, err error) {
	defer utils.TimeTrack(time.Now(), "read in cache")

	rawVal, err := r.Client.Get(context.Background(), id).Result()
	if err == redis.Nil || err != nil {
		return
	}
	t = &model.ToDoItem{}
	err = json.Unmarshal([]byte(rawVal), t)
	if err != nil {
		return
	}
	return
}

func (r *redisClient) UpdateTodoItem(t *model.ToDoItem) error {
	defer utils.TimeTrack(time.Now(), "update in cache")

	entry, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return r.Client.Set(context.Background(), t.ID.Hex(), entry, 0).Err()
}

func (r *redisClient) DeleteTodoItem(id string) error {
	defer utils.TimeTrack(time.Now(), "delete in cache")

	return r.Client.Del(context.Background(), id).Err()
}

func (r *redisClient) Ping() error {
	defer utils.TimeTrack(time.Now(), "ping redis Client")

	return r.Client.Ping(context.Background()).Err()
}

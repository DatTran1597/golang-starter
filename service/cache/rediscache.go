package cache

import (
	"time"

	"github.com/DatTran1597/golang-starter/model"
	"github.com/go-redis/redis/v7"
)

type RedisCacher struct {
	Client  redis.Client
	TimeOut time.Duration
}

func NewRedisCacher(setting *model.CacheSetting) (Cacher, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     setting.URI,
		Password: setting.Password,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	timeOut, _ := time.ParseDuration(setting.Timeout)
	return &RedisCacher{
		Client:  *redisClient,
		TimeOut: timeOut,
	}, nil
}

func (r *RedisCacher) Get(key string) (interface{}, error) {
	val, err := r.Client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *RedisCacher) Set(key string, val interface{}) error {
	err := r.Client.Set(key, val, r.TimeOut).Err()
	if err != nil {
		return err
	}

	return nil
}

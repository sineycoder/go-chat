package redisCli

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisCli struct {
	Redis redis.Conn
}

func NewRedisCli(network, host string) (r *RedisCli, err error) {
	conn, err := redis.Dial(network, host)
	if err != nil {
		fmt.Println("redis connect failed, err:", err)
		return
	}
	r = &RedisCli{Redis: conn}
	fmt.Println("redis connected...")
	return r, nil
}

func (r *RedisCli) PutString(key, value string) {
	_, err := r.Redis.Do("Set", key, value)
	if err != nil {
		fmt.Println("redis put failed, err:", err)
		return
	}
}

func (r *RedisCli) GetString(key string) (s string, err error) {
	rec, err := r.Redis.Do("Get", key)
	if err != nil {
		fmt.Println("redis get failed, err:", err)
		return
	}
	if rec == nil {
		return
	}
	return string(rec.([]byte)), nil
}

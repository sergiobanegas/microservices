package pkg

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type RedisRepository struct {
	port int
}

func NewRedisRepository(port int) *RedisRepository {
	return &RedisRepository{port}
}

func (repo *RedisRepository) newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":"+strconv.Itoa(repo.port))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func (repo *RedisRepository) Get(key string) ([]byte, error) {
	conn := repo.newPool().Get()
	res, err := redis.Bytes(redis.Bytes(conn.Do("GET", key)))
	if err == redis.ErrNil {
		return make([]byte, 0), nil
	}
	return res, err
}

func (repo *RedisRepository) Save(key string, value string) ([]byte, error) {
	conn := repo.newPool().Get()
	res, err := redis.Bytes(redis.Bytes(conn.Do("SET", key, value)))
	if err != nil {
		return nil, err
	}
	_, err = conn.Do("EXPIRE", key, 1800)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *RedisRepository) Delete(key string) error {
	conn := repo.newPool().Get()
	_, err := conn.Do("DEL", key)
	return err
}

package infrastructure

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisHandler struct {
	// connection redis.Conn
	pool *redis.Pool
}

func NewRedisHandler() *RedisHandler {
	// connection, err := redis.Dial("tcp", "redis:6379")
	// if err != nil {
	// 	return nil
	// }
	redisHandler := RedisHandler {
		pool: &redis.Pool {
			MaxIdle: 3,
			MaxActive: 0,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) { return redis.Dial("tcp", "redis:6379") },
		},
	}
	return &redisHandler
}

func (handler *RedisHandler) Set(key string, value string) error {
	conn := handler.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) Get(key string) (string, error) {
	conn := handler.pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if  err != nil {
		return value, err
	}
	return value, nil
}

func (handler *RedisHandler) Del(key string) error {
	conn := handler.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) MultiDel(keys []string) error {
	conn := handler.pool.Get()
	defer conn.Close()

	err := conn.Send("MULTI")
	if err != nil {
		return err
	}
	for _, key := range keys {
		err = conn.Send("DEL", key)
		if err != nil {
			return err
		}
	}
	_, err = conn.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) ExpireSetKey(key string, value string, second int) error {
	conn := handler.pool.Get()
	defer conn.Close()

	err := conn.Send("MULTI")
	if err != nil {
		return err
	}
	err = conn.Send("SET", key, value)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, second)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) ExpireKey(key string, second int) error {
	conn := handler.pool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, second)
	if err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) GetTtl(key string) (int, error) {
	conn := handler.pool.Get()
	defer conn.Close()

	ttl, err := redis.Int(conn.Do("TTL", key))
	if err != nil {
		return -1, err
	}
	return ttl, nil
}

func (handler *RedisHandler) RPush(key string, values []string) error {
	conn := handler.pool.Get()
	defer conn.Close()

	for _, value := range values {
		_, err := conn.Do("RPUSH", key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (handler *RedisHandler) LPop(key string, number int) ([]string, error) {
	conn := handler.pool.Get()
	defer conn.Close()

	result := []string{}
	for i := 0; i < number; i++ {
		value, err := redis.String(conn.Do("LPOP", key))	
		if err != nil {
			return result, err
		}
		result = append(result, value)
	}
	return result, nil
}

func (handler *RedisHandler) GetKeys(key string) ([]string, error) {
	conn := handler.pool.Get()
	defer conn.Close()

	keys, err := redis.Values(conn.Do("keys", key))
	if err != nil {
		return nil, err
	}
	result := []string{}
	for _, key := range keys {
		str := ""
		for _, c := range key.([]uint8) {
			str += string(c)
		}
		result = append(result, str)
	}
	return result, nil
}

func (handler *RedisHandler) Incr(key string) (int, error) {
	conn := handler.pool.Get()
	defer conn.Close()

	value, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		return -1, err
	}
	return value, nil
}
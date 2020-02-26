package infrastructure_test

import (
	inf "github.com/kindaidensan/UMR/infrastructure"
	"testing"
)

func TestNewRedisHandler(t *testing.T) {
	redisHandler := inf.NewRedisHandler()	
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
}

func TestRedisSet(t *testing.T) {
	redisHandler := inf.NewRedisHandler()	
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	err := redisHandler.Set("test", "test-value") 
	if  err != nil {
		t.Errorf("faild: set value")
		t.Errorf(err.Error())
	}
}

func TestRedisGet(t *testing.T) {
	redisHandler := inf.NewRedisHandler()	
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	err := redisHandler.Set("test-get", "test-value") 
	if  err != nil {
		t.Errorf("faild: set value")
		t.Errorf(err.Error())
	}
	value, err := redisHandler.Get("test-get")
	if err != nil {
		t.Errorf("faild: set value")
		t.Errorf(err.Error())
	}
	if value != "test-value" {
		t.Errorf("faild: incorrect value")
	}
}

func TestRedisRPush(t *testing.T) {
	redisHandler := inf.NewRedisHandler()	
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	err := redisHandler.RPush("list", []string{"1", "2"})
	if  err != nil {
		t.Errorf("faild: rpush value")
		t.Errorf(err.Error())
	}
}


func TestRedisLPop(t *testing.T) {
	redisHandler := inf.NewRedisHandler()	
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	err := redisHandler.RPush("list", []string{"1", "2", "3", "4", "5"})
	if  err != nil {
		t.Errorf("faild: rpush value")
		t.Errorf(err.Error())
	}
	values, err := redisHandler.LPop("list", 2)
	if  err != nil {
		t.Errorf("faild: lpop")
		t.Errorf(err.Error())
	}
	if len(values) != 2 {
		t.Errorf("faild: array size must be 2")

	} 
}

func TestRedisExpireSetKey(t *testing.T) {
	redisHandler := inf.NewRedisHandler()
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	err := redisHandler.ExpireSetKey("test", "value", 20)
	if err != nil {
		t.Errorf("faild: expire set")
		t.Errorf(err.Error())
	}
}

func TestRedisGetTtl(t *testing.T) {
	redisHandler := inf.NewRedisHandler()
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	err := redisHandler.ExpireSetKey("test", "value", 20)
	if err != nil {
		t.Errorf("faild: expire set")
		t.Errorf(err.Error())
	}
	ttl, err := redisHandler.GetTtl("test")
	if err != nil {
		t.Errorf("faild: get ttl")
		t.Errorf(err.Error())
	}	
	t.Log(ttl)
}

func TestGetKeys(t *testing.T) {
	redisHandler := inf.NewRedisHandler()
	if redisHandler == nil {
		t.Errorf("faild: cant create RedisHandler instance")
	}
	keys, err := redisHandler.GetKeys("*")
	if err != nil {
		t.Errorf("faild: get keys")
		t.Errorf(err.Error())
	}
	t.Log(keys)
}
package cache

import (
	"encoding/json"
	"errors"
	"facegram_file_server/config"
	"facegram_file_server/model/utilitymodel"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var redisConnection *redis.Client

func init() {
	_, err := GetRedisConnection()
	if err == nil {
		c := color.New(color.FgGreen).PrintfFunc()
		c("Successfully connected to redis\n")
	} else {
		c := color.New(color.FgRed).PrintfFunc()
		c("Failed Initializing redis Connection!\n")
	}
}

func GetRedisConnection() (*redis.Client, error) {

	if redisConnection != nil {
		return redisConnection, nil
	}

	var cfg = config.GetRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address + cfg.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ping := rdb.Ping()

	if ping.Val() != "PONG" {
		return nil, errors.New("cannot create connection to redis")
	}

	redisConnection = rdb
	return redisConnection, nil
}

//GetClientConnection func
func GetClientConnection(key string) (int, error) {
	rdb, err := GetRedisConnection()
	if err != nil {
		return 0, err
	}
	data, e0 := rdb.Get(key).Result()
	if e0 == redis.Nil {
		return 0, nil
	} else if e0 != nil {
		return 0, e0
	} else {
		number, e1 := strconv.Atoi(data)
		if e1 != nil {
			return 0, e1
		}
		return number, nil
	}
}

//SetClientConnection func
func SetClientConnection(key string, v int, ex time.Duration) error {
	rdb, err := GetRedisConnection()
	if err != nil {
		return err
	}
	err = rdb.Set(key, v, ex).Err()
	if err != nil {
		return err
	}
	return nil
}

//GetKeyTTl func
func GetKeyTTl(key string) (time.Duration, error) {
	rdb, err := GetRedisConnection()
	if err != nil {
		return 0, err
	}
	t, err := rdb.PTTL(key).Result()
	if err != nil {
		return 0, err
	}
	return t, nil
}

//GetSessionInfo func
func GetSessionInfo(key string) (string, error) {
	rdb, err := GetRedisConnection()
	if err != nil {
		return "", err
	}
	data, e0 := rdb.Get(key).Result()
	if e0 == redis.Nil {
		return "", nil
	} else if e0 != nil {
		return "", e0
	} else {
		return data, nil
	}
}

//SetSessionData func
func SetSessionData(key string, v string) error {
	rdb, err := GetRedisConnection()
	if err != nil {
		return err
	}
	err = rdb.Set(key, v, time.Duration(30*24*time.Hour)).Err()
	if err != nil {
		return err
	}
	return nil
}

//RemoveSession func
func RemoveSession(key string) (bool, error) {
	rdb, err := GetRedisConnection()
	if err != nil {
		return false, err
	}
	e0 := rdb.Del(key).Err()
	if e0 != nil {
		return false, e0
	} else {
		return true, nil
	}
}

//GetUserBlock func
func GetUserBlock(key string) (bool, *utilitymodel.UserBlock, error) {
	var model utilitymodel.UserBlock
	rdb, err := GetRedisConnection()
	if err != nil {
		return false, nil, err
	}
	data, e0 := rdb.Get(key).Result()
	if e0 == redis.Nil {
		return false, nil, nil
	} else if e0 != nil {
		return false, nil, e0
	} else {
		e1 := json.Unmarshal([]byte(data), &model)
		if e1 != nil {
			return false, nil, e1
		}
		return true, &model, nil
	}
}

//GetStringData func
func GetStringData(key string) (string, error) {
	rdb, err := GetRedisConnection()
	if err != nil {
		return "", err
	}
	data, e0 := rdb.Get(key).Result()
	if e0 == redis.Nil {
		return "", nil
	} else if e0 != nil {
		return "", e0
	} else {
		return data, nil
	}
}

//SetStringDataByTtl func
func SetStringDataByTtl(key, value string, ttl time.Duration) error {
	rdb, err := GetRedisConnection()
	if err != nil {
		return err
	}
	_, err = rdb.Set(key, value, ttl).Result()
	if err != nil {
		return err
	}
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	} else {
		return nil
	}
}

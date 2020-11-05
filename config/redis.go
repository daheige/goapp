package config

import (
	"errors"
	"fmt"

	"github.com/daheige/tigago/gredigo"
	"github.com/daheige/tigago/setting"
	"github.com/gomodule/redigo/redis"
)

// InitRedis 初始化redis client
func InitRedis(s *setting.Setting, section string, redisName string) error {
	// 初始化redis
	redisConf := &gredigo.RedisConf{}
	err := s.ReadSection("RedisCommon", redisConf)
	if err != nil {
		return fmt.Errorf("read redis section:%s error: %s", section, err.Error())
	}

	// log.Println("redis conf: ", redisConf)

	redisConf.SetRedisPool(redisName)

	return nil
}

// GetRedisObj 从连接池中获取redis client
// 用完就需要调用redisObj.Close()释放连接，防止过多的连接导致redis连接过多
// 导致当前请求而陷入长久等待，从而redis崩溃
func GetRedisObj(name string) (redis.Conn, error) {
	conn := gredigo.GetRedisClient(name)
	if conn == nil {
		return nil, errors.New("get redis client error")
	}

	return conn, nil
}

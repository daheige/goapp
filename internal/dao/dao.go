package dao

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Dao config
type Dao struct {
	dbEngine    *gorm.DB
	redisEngine redis.Conn
}

// New create an entry.
func NewDao(engine *gorm.DB) *Dao {
	return &Dao{dbEngine: engine}
}

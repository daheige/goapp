package config

import (
	"context"
	"fmt"

	"github.com/daheige/goapp/pkg/logger"
	"github.com/daheige/tigago/mysql"
	"github.com/daheige/tigago/setting"
	"github.com/jinzhu/gorm"
)

// InitDatabase 初始化db实例client
func InitDatabase(s *setting.Setting, section string, dbFdName string) error {
	// 数据库配置
	dbConf := &mysql.DbConf{}
	err := s.ReadSection(section, dbConf)
	if err != nil {
		return fmt.Errorf("read db section:%s error: %s", section, err.Error())
	}

	err = dbConf.SetDbPool() // 建立db连接池
	if err != nil {
		return err
	}

	// 为每个db设置一个engine name
	return dbConf.SetEngineName(dbFdName)
}

// GetDB 获取db实例
func GetDB(name string) *gorm.DB {
	db, err := mysql.GetDbObj(name)
	if err != nil {
		logger.Info(context.Background(), "get db install error", map[string]interface{}{
			"trace_error": err.Error(),
			"name":        name,
		})

		return &gorm.DB{}
	}

	return db
}

// CloseAllDatabase main exit defer run.
func CloseAllDatabase() {
	mysql.CloseAllDb()
}

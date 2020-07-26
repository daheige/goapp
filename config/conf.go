package config

import (
	"time"

	"github.com/daheige/goapp/pkg/setting"
)

// app.yaml section config.
var (
	AppServerConf = &AppServerSettingS{}
)

// AppServerSettingS server config.
type AppServerSettingS struct {
	AppEnv       string
	AppDebug     bool
	GRPCPort     int
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// InitConfig 读取rpc配置文件
func InitConfig(configDir string) error {
	s, err := setting.NewSetting(configDir)
	if err != nil {
		return err
	}

	err = s.ReadSection("AppServer", &AppServerConf)
	if err != nil {
		return err
	}

	AppServerConf.ReadTimeout *= time.Second
	AppServerConf.WriteTimeout *= time.Second

	// init db
	err = InitDatabase(s, "DbDefault", "default")
	if err != nil {
		return err
	}

	// 初始化redis
	err = InitRedis(s, "RedisCommon", "default")
	if err != nil {
		return err
	}

	// init mc config.
	// err = InitMc(s, "McAddress", "default")
	// if err != nil {
	// 	return err
	// }

	return nil
}

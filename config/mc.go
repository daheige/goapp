package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/daheige/tigago/setting"
)

var mcMap = map[string]*memcache.Client{}

// InitMc init memcached client.
func InitMc(s *setting.Setting, section string, mcFdName string) error {
	var mcStr string
	err := s.ReadSection("McAddress", &mcStr)
	if err != nil {
		return fmt.Errorf("read redis section:%s error: %s", section, err.Error())
	}

	// log.Println("mc config: ", mcStr)

	if mcStr == "" {
		log.Fatalln("get mc config fail", nil)
	}

	mcList := strings.Split(mcStr, ";")
	mc := memcache.New(mcList...)
	mc.Timeout = 2 * time.Second
	mc.MaxIdleConns = 300
	mcMap[mcFdName] = mc

	return nil
}

// GetMc 获得mc client
func GetMc(name string) *memcache.Client {
	return mcMap[name]
}

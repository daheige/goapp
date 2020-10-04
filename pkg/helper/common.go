package helper

import (
	"regexp"
	"strings"
)

var (
	regAndroid = regexp.MustCompile("(a|A)ndroid|dr")
	regIos     = regexp.MustCompile("i(p|P)(hone|ad|od)|(m|M)ac")
)

// GetDeviceByUa 根据ua获取设备名称
func GetDeviceByUa(ua string) string {
	plat := "web"
	if regAndroid.MatchString(ua) {
		plat = "android"
	} else if regIos.MatchString(ua) {
		plat = "ios"
	} else if strings.Contains(ua, "grpc") {
		plat = "grpc"
	}

	return plat
}

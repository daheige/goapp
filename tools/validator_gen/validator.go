package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	validatorLogDir string
)

func init() {
	flag.StringVar(&validatorLogDir, "validator_log_dir", "./", "validator log dir")
	flag.Parse()
}

func main() {
	// 读取文件的内容
	file, err := os.Open(validatorLogDir + "/validator.log")
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	// 处理结束后关闭文件
	defer file.Close()

	arr := make([]string, 0, 20)

	// 使用bufio读取
	r := bufio.NewReader(file)
	for {
		// 以分隔符形式读取,比如此处设置的分割符是\n,则遇到\n就返回,且包括\n本身 直接返回字节数数组
		data, err := r.ReadBytes('\n')

		// 读取到末尾退出
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}

		str := string(data)
		if str != "" {
			// 打印出内容
			fmt.Printf("%v", str)
			s := strings.Split(strings.Trim(str, "\n"), "=")

			log.Println("s len = ", len(s))
			// log.Println("s = ", s)

			if len(s) == 2 {
				arr = append(arr, s[1])
			}
		}
	}

	log.Println("req validator: ", arr)
	for k := range arr {
		log.Println("val: ", arr[k])
	}

}

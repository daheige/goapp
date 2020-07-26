package main

import (
	"flag"
)

var (
	validatorLogDir string
)

func init() {
	flag.StringVar(&validatorLogDir, "validator_log_dir", "./", "validator log dir")
	flag.Parse()
}

func main() {

}

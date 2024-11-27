package main

import (
	"github.com/aserafim/pos_golang/09-APIS/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	println(config.DBDriver)

	if err != nil {
		println(err)
	}
}

package main

import (
	"fmt"

	"github.com/aserafim/pos_golang/09_APIs/configs"
)

func main() {
	cfg, _ := configs.LoadConfig("/home/aserafim/Repos/pos_golang/09 - APIs/cmd/server/")
	fmt.Println(cfg.DBDriver)
}

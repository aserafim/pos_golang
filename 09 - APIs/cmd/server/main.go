package main

import (
	"fmt"

	"github.com/aserafim/pos_golang/09_APIs/configs"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
)

func main() {
	cfg, _ := configs.LoadConfig(".")
	fmt.Println(cfg.DBDriver)

	alefe, _ := entity.NewUser("Alefe", "s@s.com", "1234")

	fmt.Println(alefe.ID)
	fmt.Println(alefe.Name)
	fmt.Println(alefe.Email)
	fmt.Println(alefe.Password)
	fmt.Println(alefe.ValidatePassword("1234"))

}

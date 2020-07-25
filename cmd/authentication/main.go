package main

import (
	"fmt"

	"github.com/PavluninVladimir/go-project-server/cmd/authentication/config"
)

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.ConfigVar)
}

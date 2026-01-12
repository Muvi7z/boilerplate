package main

import (
	"fmt"
	"github.com/Muvi7z/boilerplate/notification/internal/config"
)

const configPath = "./deploy/compose/notification/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

}

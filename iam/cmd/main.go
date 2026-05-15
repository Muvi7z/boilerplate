package main

import (
	"fmt"
	"github.com/jcmturner/gokrb5/v8/config"
)

const configPath = "../deploy/compose/inventory/.env"

func main() {
	err, _ := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config file: %s", err))
	}

}

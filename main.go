package main

import (
	"fmt"

	"github.com/FabsMS/Gopportunies/config"
	"github.com/FabsMS/Gopportunies/router"
)

var (
	logger config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Inicializando Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Inicializando o Router
	router.Initializer()
}

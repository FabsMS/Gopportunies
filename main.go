package main

import (
	"fmt"

	"github.com/FabsMS/Gopportunies/config"
	"github.com/FabsMS/Gopportunies/router"
)

func main() {

	// Inicializando Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Inicializando o Router
	router.Initializer()
}

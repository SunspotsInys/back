package main

import (
	"strconv"

	"github.com/SunspotsInys/thedoor/configs"
	"github.com/SunspotsInys/thedoor/routes"
)

func main() {

	routes.InitRoutes().Run(":" + strconv.Itoa(configs.Conf.Port))

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGINT)
	// <-c
}

package main

import (
	"underworld/routers"
	"underworld/utils"
)

func main() {
	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT", ":8080")
	router.Run(port)
}

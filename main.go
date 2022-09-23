package main

import (
	"flag"
	"odisha_gov_be/utils"
)

func main() {
	appConfig := &AppConfig{}
	flag.StringVar(&appConfig.Env, "env", "d", "environment variable to use for running application")
	flag.IntVar(&appConfig.Port, "port", 3000, "port to run application on")
	flag.Parse()

	// Initilization
	utils.InitLogger(appConfig.Env)
	utils.InitDB()

	utils.Logger().Info("initilizations complete")
	utils.Logger().Info("app configs are:", *appConfig)

	// Initilization server
	server := InitServer(appConfig)
	InitRoutes(server)

	// start server
	StartServer()
}

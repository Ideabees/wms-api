package main

import (
	"wms-app/config"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/routes"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&dbModels.User{})

	r := routes.SetupRoutes()
	r.Run(":8080")
}

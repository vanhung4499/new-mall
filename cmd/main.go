package main

import (
	"fmt"
	"new-mall/internal/config"
	"new-mall/internal/database"
	"new-mall/internal/global"
	"new-mall/internal/routes"
)

func main() {
	// Loading configuration
	global.CONFIG = config.LoadConfig()
	// Loading database
	global.DB = database.NewDatabase()
	database.Migrate()
	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	_ = r.Run(global.CONFIG.System.HttpPort)
	fmt.Println("Start successfully...")
}

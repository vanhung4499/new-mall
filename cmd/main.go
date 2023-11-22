package main

import (
	"fmt"
	"new-mall/config"
	"new-mall/pkg/utils"
	"new-mall/repository/cache"
	"new-mall/repository/db/dao"
	"new-mall/route"
)

func main() {
	loading()
	r := route.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
	fmt.Println("Start successfully...")
}

func loading() {
	config.InitConfig()
	dao.InitMySQL()
	cache.InitCache()
	utils.InitLog()
	fmt.Println("Loading configuration completed...")
	go scriptStarting()
}

func scriptStarting() {
	// start some scripts
}

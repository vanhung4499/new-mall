package core

import (
	"fmt"
	"go.uber.org/zap"
	"new-mall/global"
	initialize "new-mall/inittialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.Config.System.UseMultipoint || global.Config.System.UseRedis {
		// Initialize redis service
		initialize.Redis()
	}
	if global.Config.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// Ensure text is output in order
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Log.Info("server run success on ", zap.String("address", address))

	global.Log.Error(s.ListenAndServe().Error())
}

package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"new-mall/global"
	"new-mall/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.Config.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.Config.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}

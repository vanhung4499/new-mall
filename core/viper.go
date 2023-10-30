package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"new-mall/core/internal"
	"new-mall/global"
	"os"
)

// Viper //
// Priority: Command line > Environment variables > Default value
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("You are using the %s environment name in gin mode, and the path to config is %s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("You are using the %s environment name in gin mode, and the path to config is %\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("You are using the %s environment name in gin mode, and the path to config is %s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // The environment variable stored in internal.ConfigEnv constant is not empty and assigns the value to config
				config = configEnv
				fmt.Printf("You are using the %s environment variable, the path to config is %s\n", internal.ConfigEnv, config)
			}
		} else { // The command line parameters are not empty. Assign the value to config.
			fmt.Printf("You are using the value passed by the -c parameter on the command line, and the path to config is %s\n", config)
		}
	} else { // The first value of the variable parameter passed by the function is assigned to config
		config = path[0]
		fmt.Printf("You are using the value passed by func Viper(), the path to config is %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	return v
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokuovo/project-layout/configs"
	route "github.com/gokuovo/project-layout/pkg/router"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	configs.InitDB()
	r := gin.Default()
	route.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}

func InitConfig() {
	//workDir, _ := os.Getwd()
	//viper.SetConfigName("application-dev")
	//viper.SetConfigType("yml")
	viper.SetConfigFile("./configs/application-test.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}

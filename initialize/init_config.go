package initialize

import (
	"gin_skeleton/g"
	"github.com/spf13/viper"
)

func InitConfig() {

	//读取配置文件
	viper.AddConfigPath("./conf")
	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}
	g.Config = viper.GetViper()
}

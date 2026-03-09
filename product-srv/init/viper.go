package init

import (
	"fmt"
	"github.com/spf13/viper"
	"product-srv/config"
)

func ViperInit() {
	viper.SetConfigFile("dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	err = viper.Unmarshal(&config.Config)
	if err != nil {
		panic("配置文件解析失败")
	}
	fmt.Println("配置文件读取成功", config.Config)
}

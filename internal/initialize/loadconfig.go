package initialize

import (
	"fmt"
	"tudo-thrfting-clothes-server/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Initialize viper
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("production") // development || production
	viper.SetConfigType("yaml")

	// // Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	} else {
		fmt.Println("Config file loaded successfully")
	}

	// Get the value from the configuration file
	// fmt.Println("Name: ", viper.GetString("database.user"))
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Error unmarshal config: ", err)
	}
}

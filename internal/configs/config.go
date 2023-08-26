package config

import (
	"avesto-service/internal/models"
	"github.com/spf13/viper"
	"gopkg.in/lumberjack.v2"
	"log"
)

func InitConfig() (cfg models.Configuration) {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml") // Explicitly set the config file type
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err = viper.UnmarshalKey("configuration", &cfg) // Unmarshal the 'configuration' key
	if err != nil {
		log.Println(err)
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    2,  // megabytes
		MaxAge:     40, // days
		MaxBackups: 30,
		Compress:   true,
	})

	log.Println("---*--- Starting Logging ---*---")
	return
}

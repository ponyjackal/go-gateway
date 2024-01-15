package main

import (
	"time"

	"github.com/ponyjackal/go-gateway/internal/app/routers"
	"github.com/ponyjackal/go-gateway/pkg/config"
	"github.com/ponyjackal/go-gateway/pkg/logger"

	"github.com/spf13/viper"
)

func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}

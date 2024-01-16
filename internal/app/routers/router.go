package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ponyjackal/go-gateway/internal/app/middlewares"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
)

func SetupRoute() *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	limiter := rate.NewLimiter(1, 5)
	router.Use(middlewares.RateLimitMiddleware(limiter))

	RegisterRoutes(router) //routes register

	return router
}

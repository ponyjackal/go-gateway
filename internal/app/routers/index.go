package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ponyjackal/go-gateway/internal/app/controllers"
	"github.com/ponyjackal/go-gateway/internal/domain/services"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	httpService := services.NewHTTPService()
	podcastService := services.NewPodcastService(httpService)
	podcastController := controllers.NewPodcastController(podcastService)

	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	route.GET("/podcasts", podcastController.GetPodcasts)
}

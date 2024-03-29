package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ponyjackal/go-gateway/internal/app/controllers"
	"github.com/ponyjackal/go-gateway/internal/app/graphql"
	"github.com/ponyjackal/go-gateway/internal/domain/services"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	httpService := services.NewHTTPService()
	podcastService := services.NewPodcastService(httpService)
	podcastController := controllers.NewPodcastController(podcastService)
	graphqlService := graphql.NewGraphQLService(podcastService)

	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	route.GET("/heartbeat", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "alive"}) })

	route.GET("/podcasts", podcastController.GetPodcasts)

	// graphql
	route.POST("/graphql", graphqlService.GraphqlHandler)
	route.GET("/playground", graphqlService.GraphqlPlaygroundHandler)
}

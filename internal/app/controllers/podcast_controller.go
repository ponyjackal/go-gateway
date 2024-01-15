package controllers

import "github.com/gin-gonic/gin"

type PodcastController struct {
}

func (c *PodcastController) NewPodcastController() *PodcastController {
	return &PodcastController{}
}

func (c *PodcastController) GetPodcasts(ctx *gin.Context) {
	// Retrieve query parameters
	// page, pageExists := ctx.GetQuery("page")
	// limit, limitExists := ctx.GetQuery("limit")
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ponyjackal/go-gateway/internal/domain/services"
	"github.com/ponyjackal/go-gateway/pkg/types"
)

type PodcastController struct {
	podcastService *services.PodcastService
}

func NewPodcastController(podcastService *services.PodcastService) *PodcastController {
	return &PodcastController{
		podcastService: podcastService,
	}
}

func (c *PodcastController) GetPodcasts(ctx *gin.Context) {
	// Retrieve all query parameters as a map
	queryParams := ctx.Request.URL.Query()
	var query types.GetPodcastsQuery

	for key, values := range queryParams {
		switch key {
		case "page":
			// Retrieve the "page" query parameters (if provided)
			if pageStr := values[0]; pageStr != "" {
				p, err := strconv.Atoi(pageStr)
				if err == nil && p > 0 {
					query.Page = &p
				}
			}
		case "limit":
			// Retrieve the "limit" query parameters (if provided)
			if limitStr := values[0]; limitStr != "" {
				l, err := strconv.Atoi(limitStr)
				if err == nil && l > 0 {
					query.Limit = &l
				}
			}
		case "search":
			// Retrieve the "search" query parameters (if provided)
			query.Search = &values[0]
		case "title":
			// Retrieve the "title" query parameters (if provided)
			query.Title = &values[0]
		case "categoryName":
			// Retrieve the "categoryName" query parameters (if provided)
			query.CategoryName = &values[0]
		default:
		}
	}

	podcasts, err := c.podcastService.GetPodcasts(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, &podcasts)
}

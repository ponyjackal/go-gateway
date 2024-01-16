package graphql

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/ponyjackal/go-gateway/internal/domain/services"
	"github.com/ponyjackal/go-gateway/pkg/types"
)

type GraphQLService struct {
	podcastService *services.PodcastService
	schema         graphql.Schema
}

func NewGraphQLService(podcastService *services.PodcastService) *GraphQLService {
	service := &GraphQLService{
		podcastService: podcastService,
	}
	service.schema = service.createSchema()
	return service
}

func (s *GraphQLService) createSchema() graphql.Schema {
	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"podcasts": &graphql.Field{
				Type: graphql.NewList(podcastType),
				Args: graphql.FieldConfigArgument{
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"search": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"categoryName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.getPodcastsResolver,
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	return schema
}

func (s *GraphQLService) getPodcastsResolver(p graphql.ResolveParams) (interface{}, error) {
	var query types.GetPodcastsQuery

	// Extracting and setting the arguments
	if page, ok := p.Args["page"].(int); ok {
		query.Page = &page
	}
	if limit, ok := p.Args["limit"].(int); ok {
		query.Limit = &limit
	}
	if search, ok := p.Args["search"].(string); ok {
		query.Search = &search
	}
	if title, ok := p.Args["title"].(string); ok {
		query.Title = &title
	}
	if categoryName, ok := p.Args["categoryName"].(string); ok {
		query.CategoryName = &categoryName
	}

	// Call the podcast service to get podcasts
	podcasts, err := s.podcastService.GetPodcasts(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching podcasts: %v", err)
	}

	return podcasts, nil
}

func (s *GraphQLService) GraphqlHandler(c *gin.Context) {
	// Extract query from request
	var req GraphqlRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute GraphQL query using the schema
	result := graphql.Do(graphql.Params{
		Schema:         s.schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
		OperationName:  req.OperationName,
	})
	if len(result.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, result.Errors)
		return
	}

	// Send the response
	c.JSON(http.StatusOK, result)
}

const graphqlPlaygroundHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset=utf-8/>
    <title>GraphQL Playground</title>
    <link href="https://cdn.jsdelivr.net/npm/graphql-playground-react@1.7.20/build/static/css/index.css" rel="stylesheet" />
    <script src="https://cdn.jsdelivr.net/npm/graphql-playground-react@1.7.20/build/static/js/middleware.js"></script>
</head>
<body>
    <div id="root"></div>
    <script>window.addEventListener('load', function(event) {
        GraphQLPlayground.init(document.getElementById('root'), {
            endpoint: '/graphql' // Update this endpoint to your GraphQL endpoint
        })
    })</script>
</body>
</html>
`

func (s *GraphQLService) GraphqlPlaygroundHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(graphqlPlaygroundHTML))
}

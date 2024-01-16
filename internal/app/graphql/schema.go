package graphql

import "github.com/graphql-go/graphql"

type GraphqlRequestBody struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
}

var podcastImagesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PodcastImages",
	Fields: graphql.Fields{
		"default": &graphql.Field{
			Type: graphql.String,
		},
		"featured": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
		"wide": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var podcastType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Podcast",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"images": &graphql.Field{
			Type: podcastImagesType,
		},
		"isExclusive": &graphql.Field{
			Type: graphql.Boolean,
		},
		"publisherName": &graphql.Field{
			Type: graphql.String,
		},
		"publisherId": &graphql.Field{
			Type: graphql.String,
		},
		"mediaType": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"categoryId": &graphql.Field{
			Type: graphql.String,
		},
		"categoryName": &graphql.Field{
			Type: graphql.String,
		},
		"hasFreeEpisodes": &graphql.Field{
			Type: graphql.Boolean,
		},
		"playSequence": &graphql.Field{
			Type: graphql.String,
		},
	},
})

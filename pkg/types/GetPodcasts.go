package types

type GetPodcastsQuery struct {
	Page         *int    `json:"page,omitempty"`
	Limit        *int    `json:"limit,omitempty"`
	Search       *string `json:"search,omitempty"`
	Title        *string `json:"title,omitempty"`
	CategoryName *string `json:"categoryName,omitempty"`
}

type Podcast struct {
	ID              string        `json:"id"`
	Title           string        `json:"title"`
	Images          PodcastImages `json:"images"`
	IsExclusive     bool          `json:"isExclusive"`
	PublisherName   string        `json:"publisherName"`
	PublisherID     string        `json:"publisherId"`
	MediaType       string        `json:"mediaType"`
	Description     string        `json:"description"`
	CategoryID      string        `json:"categoryId"`
	CategoryName    string        `json:"categoryName"`
	HasFreeEpisodes bool          `json:"hasFreeEpisodes"`
	PlaySequence    string        `json:"playSequence"`
}

type PodcastImages struct {
	Default   string `json:"default"`
	Featured  string `json:"featured"`
	Thumbnail string `json:"thumbnail"`
	Wide      string `json:"wide"`
}

type GetPodcastsResponse struct {
	Podcasts []Podcast `json:"podcasts"`
}

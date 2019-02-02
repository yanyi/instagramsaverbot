package scraper

type postService interface {
	getOnlyImageURLs() error
	appendImageURLs(shortCodeMedia) error
}

type post struct {
	post      postService
	GraphQL   graphQL `json:"graphql"`
	ImageURLs []string
}

type graphQL struct {
	ShortCodeMedia shortCodeMedia `json:"shortcode_media"`
}

// shortCodeMedia represents the required information for image URLs.
type shortCodeMedia struct {
	DisplayResources      []displayResources    `json:"display_resources"`
	IsVideo               bool                  `json:"is_video"`
	EdgeSidecarToChildren edgeSidecarToChildren `json:"edge_sidecar_to_children"`
}

type displayResources struct {
	Src          string `json:"src"`
	ConfigWidth  int32  `json:"config_width"`
	ConfigHeight int32  `json:"config_height"`
}

type edgeSidecarToChildren struct {
	Edges []edges `json:"edges"`
}

type edges struct {
	Node shortCodeMedia `json:"node"`
}

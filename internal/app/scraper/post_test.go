package scraper

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PostSuite struct {
	suite.Suite
	mixedPost post
	badPost   post
	goodPost  post
}

func TestPostSuite(t *testing.T) {
	suite.Run(t, new(PostSuite))
}

func (suite *PostSuite) SetupTest() {
	suite.mixedPost = post{
		GraphQL: graphQL{
			shortCodeMedia{
				EdgeSidecarToChildren: edgeSidecarToChildren{
					Edges: []edges{
						edges{
							Node: shortCodeMedia{
								DisplayResources: []displayResources{
									displayResources{
										Src:          "https://photo-small/",
										ConfigWidth:  400,
										ConfigHeight: 400,
									},
									displayResources{
										Src:          "https://photo-large/",
										ConfigWidth:  1080,
										ConfigHeight: 1080,
									},
								},
								IsVideo: false,
							},
						},
						edges{
							Node: shortCodeMedia{
								DisplayResources: []displayResources{
									displayResources{
										Src: "https://video/",
									},
								},
								IsVideo: true,
							},
						},
					},
				},
			},
		},
	}

	suite.goodPost = post{
		GraphQL: graphQL{
			shortCodeMedia{
				DisplayResources: []displayResources{
					displayResources{
						Src:          "https://photo-small/",
						ConfigWidth:  400,
						ConfigHeight: 400,
					},
					displayResources{
						Src:          "https://photo-large/",
						ConfigWidth:  1080,
						ConfigHeight: 1080,
					},
				},
			},
		},
	}

	suite.badPost = post{
		GraphQL: graphQL{
			shortCodeMedia{},
		},
	}
}

// Should get only image URLs from a post with 1 image, 1 video
func (suite *PostSuite) TestGetOnlyImageURLs() {
	post := suite.mixedPost
	post.getOnlyImageURLs()
	suite.Equal(len(post.ImageURLs), 1)
}

func (suite *PostSuite) TestAppendImageURLs() {
	scm := suite.goodPost.GraphQL.ShortCodeMedia
	post := suite.goodPost
	err := post.appendImageURLs(scm)
	suite.NoError(err)
	suite.Equal(len(post.ImageURLs), 1)
}

// Should throw error if can't get highest definition image
func (suite *PostSuite) TestGetHighestDefImageURLFromBadPost() {
	scm := suite.badPost.GraphQL.ShortCodeMedia
	_, err := scm.getHighestDefImageURL()
	suite.Error(err)
}

// Should get highest definition image when post is ok
func (suite *PostSuite) TestGetHighestDefImageURLFromGoodPost() {
	scm := suite.goodPost.GraphQL.ShortCodeMedia
	imageURL, err := scm.getHighestDefImageURL()
	suite.Exactly("https://photo-large/", imageURL)
	suite.NoError(err)
}

package scraper

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ScraperSuite struct {
	suite.Suite
}

func TestScraperSuite(t *testing.T) {
	suite.Run(t, new(ScraperSuite))
}

func (suite *ScraperSuite) TestPopulateInstagramURLs() {
	// Might need to fake/mock HTTP response here
	// populateInstagramURLs(*http.Response, *[]string) error
}

// func (suite *ScraperSuite) TestUnmarshalJSON() {
// 	// unmarshalJSON([]byte, *post)
// }

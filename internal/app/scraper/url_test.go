package scraper

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	errMsg = "Link given is not an Instagram link"
)

type URLHandlersSuite struct {
	suite.Suite
}

func TestURLHandlersSuite(t *testing.T) {
	suite.Run(t, new(URLHandlersSuite))
}

// Should throw error when given bad URL scheme
func (suite *URLHandlersSuite) TestGetGraphURLBad() {
	_, err := getGraphURL(":url")
	suite.Error(err)
	suite.EqualError(err, errMsg)
}

// Should throw error when given non-Instagram link
func (suite *URLHandlersSuite) TestGetGraphURLNotInstagram() {
	_, err := getGraphURL("https://google.com/")
	suite.Error(err)
	suite.EqualError(err, errMsg)
}

// Should append "?__a=1" at back of Instagram link
func (suite *URLHandlersSuite) TestGetGraphURL() {
	u, err := getGraphURL("https://instagram.com/p/random/")
	suite.NoError(err)
	suite.Equal("https://instagram.com/p/random/?__a=1", u)
}

// Should pass when given good URL scheme
func (suite *URLHandlersSuite) TestGetRawURLGood() {
	url, _ := url.Parse("https://test/?a=bc")
	u := getRawURL(url)
	suite.Exactly("https://test/", u)
}

// Should give error if it's not an Instagram URL
func (suite *URLHandlersSuite) TestCheckIfInstagram() {
	_, errorGoogle := urlIsInstagram("https://google.com/")
	suite.Error(errorGoogle)
	suite.EqualError(errorGoogle, errMsg)

	instagram, errorInstagram := urlIsInstagram("https://instagram.com/")
	suite.NoError(errorInstagram)
	suite.Equal("instagram.com", instagram.Hostname())
}

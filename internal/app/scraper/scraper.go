// Package scraper does the scraping of Instagram links.
//It will return an error immediately if the link is not an Instagram link.
package scraper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	expectedResponseType = "application/json; charset=utf-8"
)

// Scrape scrapes the Instagram URL.
// Takes in a string input (URL) and a slice to be appended.
// It will return an error if the given URL is not an Instagram link, or the
// given Instagram link does not correspond to a post with response.
func Scrape(inputURL string, urls *[]string) error {
	url, err := getGraphURL(inputURL)
	if err != nil {
		// Throw this error back up to the user so that they can resubmit
		return err
	}

	req := getRequest(url)
	client := &http.Client{}
	res := getResponse(client, req)
	defer res.Body.Close()

	err = populateInstagramURLs(res, urls)
	if err != nil {
		return err
	}

	// No image found
	if len(*urls) == 0 {
		return errors.New("No image found in the Instagram post")
	}

	return nil
}

// populateInstagramURLs will take in the HTTP Response and a URLs slice, and then append to the URLs slice.
func populateInstagramURLs(res *http.Response, urls *[]string) error {
	var post post
	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseContentType := res.Header["Content-Type"][0]
	notEmptyResponse := len(resBodyBytes) > 0

	if notEmptyResponse && responseContentType == expectedResponseType {
		unmarshalJSON(resBodyBytes, &post)
		err := post.getOnlyImageURLs()
		if err != nil {
			return err
		}
	} else {
		return errors.New("Unable to process link given. Please try again")
	}

	*urls = post.ImageURLs

	return nil
}

// unmarshalJSON is a wrapper for json.Unmarshal().
func unmarshalJSON(resBodyBytes []byte, post *post) {
	if err := json.Unmarshal(resBodyBytes, &post); err != nil {
		log.Fatal("unmarshalJSON()", err)
	}
}

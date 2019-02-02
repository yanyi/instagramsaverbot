package scraper

import (
	"log"
	"net/http"
)

func getRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("getRequest() ", err)
	}

	return req
}

func getResponse(client *http.Client, request *http.Request) *http.Response {
	res, err := client.Do(request)
	if err != nil {
		log.Fatal("getResponse() ", err)
	}

	return res
}

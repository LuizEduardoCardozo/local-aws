package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type AwsRequest struct {
	region  string
	service string
}

func GetAuthHeader(res http.ResponseWriter, req *http.Request) {
	header := req.Header["Authorization"][0]
	response := ParseAuthHeader(header)
	fmt.Printf("region: %s, service: %s\n", response.region, response.service)
}

func ParseAuthHeader(authHeader string) AwsRequest {
	regexMatch := regexp.MustCompile(`20\d*\/(?P<region>\w*-\w*-\d)\/(?P<service>\w*)\/`)
	match := regexMatch.FindStringSubmatch(authHeader)
	return AwsRequest{
		region:  match[1],
		service: match[2],
	}
}

func main() {
	http.HandleFunc("/", GetAuthHeader)
	http.ListenAndServe(":3000", nil)
}


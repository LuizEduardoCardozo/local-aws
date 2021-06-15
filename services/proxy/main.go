package main

import (
	"fmt"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on port %s\n", port)
	http.HandleFunc("/", GetAuthHeader)
	http.ListenAndServe(addr, nil)
}

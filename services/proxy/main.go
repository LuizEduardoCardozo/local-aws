package main

import (
	"fmt"
	"laws/lib"
	"net/http"
	"os"
	"regexp"
)

func ParseAuthHeader(authHeader string) (string, string) {
	regexMatch := regexp.MustCompile(`20\d*\/(?P<region>\w*-\w*-\d)\/(?P<service>\w*)\/`)
	match := regexMatch.FindStringSubmatch(authHeader)
	return match[1], match[2]
}

func GetAuthHeader(res http.ResponseWriter, req *http.Request) {
	header := req.Header["Authorization"][0]
	region, service := ParseAuthHeader(header)
	awsrequest := lib.AwsRequest{
		Region:  region,
		Service: service,
		Req:     req,
	}
	switch awsrequest.Service {
	case "s3":
		lib.S3Parse(awsrequest)
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

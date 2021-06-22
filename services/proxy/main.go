package main

import (
	"errors"
	"fmt"
	"laws/lib"
	"laws/lib/interfaces"
	"laws/lib/parsers"
	"net/http"
	"os"
	"regexp"
)

func _ParseAuthHeader(authHeader string) (string, string) {
	regexMatch := regexp.MustCompile(`20\d*\/(?P<region>\w*-\w*-\d)\/(?P<service>\w*)\/`)
	match := regexMatch.FindStringSubmatch(authHeader)
	return match[1], match[2]
}

func _GetAWSParser(awsRequest lib.AwsRequest) (interfaces.AWSParser, error) {
	parsersMap := map[string]interfaces.AWSParser{
		"s3":  parsers.S3Parser{},
		"sqs": parsers.SQSParser{},
	}

	if awsParser, keyExists := parsersMap[awsRequest.Service]; keyExists {
		return awsParser, nil
	}

	return nil, errors.New("Service not found: " + awsRequest.Service)
}

func _GetAWSRequest(req *http.Request) lib.AwsRequest {
	header := req.Header["Authorization"][0]
	region, service := _ParseAuthHeader(header)
	return lib.AwsRequest{
		Region:  region,
		Service: service,
		Req:     req,
	}
}

func _HandleRequest(res http.ResponseWriter, req *http.Request) {
	awsRequest := _GetAWSRequest(req)
	parser, err := _GetAWSParser(awsRequest)
	if err != nil {
		fmt.Println("Deu caca", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	parser.Parse(awsRequest)
	res.WriteHeader(http.StatusOK)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on port %s\n", port)

	http.HandleFunc("/", _HandleRequest)
	http.ListenAndServe(addr, nil)
}

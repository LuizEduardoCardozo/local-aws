package lib

import "net/http"

type AwsRequest struct {
	Region  string
	Service string
	Req     *http.Request
}

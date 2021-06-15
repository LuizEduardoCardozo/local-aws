package lib

import "fmt"

func S3Parse(awsreq AwsRequest) {
	fmt.Printf("service: %s region: %s\n", awsreq.Service, awsreq.Region)
}

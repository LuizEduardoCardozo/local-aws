package parsers

import (
	"fmt"
	"laws/lib"
)

type SQSParser struct{}

func (sqsParser SQSParser) Parse(awsRequest lib.AwsRequest) {
	fmt.Println("parsing SQS", awsRequest)
}

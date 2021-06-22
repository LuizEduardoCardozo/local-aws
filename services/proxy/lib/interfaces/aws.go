package interfaces

import "laws/lib"

// Base AWSParser interface
type AWSParser interface {
	Parse(awsRequest lib.AwsRequest)
}

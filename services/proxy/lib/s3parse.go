package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

type BucketData struct {
	BucketName string
	FilePath   string
}

func S3Parse(awsreq AwsRequest) {
	fmt.Printf("Method: %s\n", awsreq.Req.Method)
	fmt.Printf("URL: %s\n", awsreq.Req.URL.Path)

	url := awsreq.Req.URL.Path
	regexMatch := regexp.MustCompile(`\/(?P<bucket_name>\w*)\/|(?P<file_path>.*)`)
	match := regexMatch.FindAllStringSubmatch(url, 2)

	var instructions BucketData

	if len(match) == 2 {
		instructions = BucketData{
			BucketName: match[0][0],
			FilePath:   match[1][0],
		}
	} else {
		instructions = BucketData{
			BucketName: match[0][0],
			FilePath:   "",
		}
	}
	fmt.Printf("%+v%c", instructions, 0xA)

	switch awsreq.Req.Method {
	case "GET":
		fmt.Println("Listar buckets")
	case "PUT":
		if len(match) == 1 {
			fmt.Println("Criar novo bucket")
		} else {
			fmt.Println("Upload de arquivo para o bucket")
			fileBytes, err := ioutil.ReadAll(awsreq.Req.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("Request body: %s\n", fileBytes)
		}
	}
}

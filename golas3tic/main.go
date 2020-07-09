package main

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	s3pkg "github.com/davidthorpe71/go-elastic/golas3tic/s3Part"
)

type testXML struct {
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

var testXML1 = `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<david>\n<id>id123</id>\n<title>DaveyTitle</title>\n<description>Davey description</description>\n</david>`

func main() {
	s3Uploader := s3pkg.CreateUploader()

	// xmlTest := &testXML{
	// 	ID:          "1234",
	// 	Title:       "First title",
	// 	Description: "second test description",
	// }

	content := []byte(testXML1)

	params := &s3.PutObjectInput{
		Bucket:      aws.String("golas3tic-test-bucket"),
		Key:         aws.String("fourthFile.xml"),
		ContentType: aws.String("text/plain"),
		Body:        bytes.NewReader(content),
	}

	result, err2 := s3Uploader.PutObject(params)

	if err2 != nil {
		fmt.Println("S3 upload error:", err2)
	}

	fmt.Println("Results:", result)

}

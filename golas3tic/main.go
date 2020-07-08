package main

import (
	"bytes"
	"encoding/xml"
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

const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

func main() {
	s3Uploader := s3pkg.CreateUploader()

	xmlTest := &testXML{
		ID:          "1234",
		Title:       "First title",
		Description: "second test description",
	}

	file, err := xml.MarshalIndent(xmlTest, "", " ")
	file = []byte(xml.Header + string(file))
	if err != nil {
		fmt.Println("Couldn't marshal string into xml")
	}

	reader := bytes.NewReader(file)
	upParams := &s3.PutObjectInput{
		Bucket: aws.String("golas3tic-test-bucket"),
		Key:    aws.String("thirdFile.xml"),
		Body:   reader,
	}

	result, err2 := s3Uploader.PutObject(upParams)

	if err2 != nil {
		fmt.Println("S3 upload error:", err2)
	}

	fmt.Println("Results:", result)

}

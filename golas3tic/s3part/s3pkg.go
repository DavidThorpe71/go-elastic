package s3pkg

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type s3ClientInterfcae interface {
	CreateUploader() *s3.S3
	Put() error
	// Get() error
}

type s3Client struct {
	s3      s3iface.S3API
	session *session.Session
}

// CreateSession creates a new aws session
func (s3c *s3Client) NewSession() *session.Session {
	//  Start new session for s3 upload
	// sess := session.Must(session.NewSession(&aws.Config{
	// 	Region:      aws.String("eu-west-2"),
	// 	Credentials: credentials.NewSharedCredentials("", "golas3tic"),
	// }))
	region := endpoints.EuWest2RegionID
	credentials := credentials.NewSharedCredentials("", "golas3tic")

	// S3 client for uploading
	s3Svc := session.New(aws.NewConfig().WithRegion(region).WithCredentials(credentials))

	return s3Svc
}

func (s3c *s3Client) Put(path, content string) error {
	contentBytes := []byte(content)

	params := &s3.PutObjectInput{
		Bucket:      aws.String("golas3tic-test-bucket"),
		Key:         aws.String("fourthFile.xml"),
		ContentType: aws.String("text/plain"),
		Body:        bytes.NewReader(contentBytes),
	}

	result, err2 := s3c.s3.PutObject(params)

	if err2 != nil {
		fmt.Println("S3 upload error:", err2)
	}
}

package s3pkg

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CreateSession creates a new aws session
func CreateUploader() *s3.S3 {

	//  Start new session for s3 upload
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewSharedCredentials("", "golas3tic"),
	}))

	// S3 client for uploading
	s3Svc := s3.New(sess)

	return s3Svc
}

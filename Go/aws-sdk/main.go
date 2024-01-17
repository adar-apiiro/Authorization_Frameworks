package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/aws/session"
)

func main() {
	// Replace these values with your own
	roleARN := "arn:aws:iam::ACCOUNT_ID:role/ROLE_NAME"
	region := "us-west-2"

	// Create a new AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	// Assume the IAM role
	creds := stscreds.NewCredentials(sess, roleARN)
	credentials, err := creds.Get(context.TODO())
	if err != nil {
		fmt.Println("Error assuming role:", err)
		return
	}

	// Use the assumed credentials to make an AWS service call
	// For example, you can create an S3 client with the assumed credentials
	s3Client := NewS3Client(credentials)

	// Now you can use s3Client to interact with AWS services as the assumed role

	// Example: List buckets in S3
	buckets, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Error listing S3 buckets:", err)
		return
	}

	fmt.Println("S3 Buckets:")
	for _, bucket := range buckets.Buckets {
		fmt.Println(*bucket.Name)
	}
}

// NewS3Client creates a new S3 client with the provided credentials
func NewS3Client(credentials aws.CredentialsProvider) *s3.Client {
	cfg := aws.Config{
		Credentials: credentials,
	}

	return s3.NewFromConfig(cfg)
}

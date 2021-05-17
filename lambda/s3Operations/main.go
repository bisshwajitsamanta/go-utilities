package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var s3session *s3.S3

const (
	BUCKET_NAME = "bishwajitsamanta9861"
	REGION      = "us-east-1"
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func ListBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

func CreateBucket() (resp *s3.CreateBucketOutput) {

	resp, err := s3session.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(BUCKET_NAME),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println("Bucket already Exists")
				os.Exit(1)
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println("You already owned this Bucket!")
			default:
				panic(err)
			}
		}
	}
	return resp
}

func main() {
	fmt.Println(ListBuckets())
	fmt.Println(CreateBucket())
	fmt.Println(ListBuckets())
}

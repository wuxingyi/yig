package _go

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func MakeBucket(bucketName string) (err error) {
	s3client := NewS3()
	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}
	if _, err = s3client.CreateBucket(params); err != nil {
		return err
	}
	return
}

func DeleteBucket(bucketName string) (err error) {
	s3client := NewS3()
	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	if _, err = s3client.DeleteBucket(params); err != nil {
		return err
	}
	return
}

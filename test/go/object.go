package _go

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
)

func (s3client *S3Client) PutObject(bucketName, key, value string) (err error) {
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader([]byte(value)),
	}
	if _, err = s3client.Client.PutObject(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) HeadObject(bucketName, key string) (err error) {
	params := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	_, err = s3client.Client.HeadObject(params)
	if err != nil {
		return err
	}
	return
}

func (s3client *S3Client) GetObject(bucketName, key string) (value string, err error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	out, err := s3client.Client.GetObject(params)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(out.Body)
	return string(data), err
}

func (s3client *S3Client) DeleteObject(bucketName, key string) (err error) {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	_, err = s3client.Client.DeleteObject(params)
	if err != nil {
		return err
	}
	return
}

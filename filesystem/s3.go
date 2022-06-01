package filesystem

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

type S3Drive struct {
	config   *S3DriveConfig
	s3client *s3.Client
}

type S3DriveConfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Region    string
	Endpoint  string
}

var _ Drive = (*S3Drive)(nil)

func NewS3Drive(config *S3DriveConfig) Drive {
	drive := &S3Drive{
		config: config,
	}

	drive.s3client = drive.newS3Client()

	return drive
}

func (s *S3Drive) newS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(s.config.Region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     s.config.AccessKey,
				SecretAccessKey: s.config.SecretKey,
				Source:          "provider",
			},
		}),
	)

	if err != nil {
		panic(err)
	}

	return s3.NewFromConfig(cfg)
}

func (s *S3Drive) Exists(path string) bool {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Get(path string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Put(path, contents string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Prepend(path, contents string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Append(path, contents string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Delete(path string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Copy(from, to string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Move(from, to string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Rename(from, to string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Size(path string) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) LastModified(path string) (time.Time, error) {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) Files(directory string) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func (s *S3Drive) AllFiles(directory string) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

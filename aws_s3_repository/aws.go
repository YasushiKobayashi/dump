package aws_s3_repository

import (
	"bytes"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

type (
	AwsS3Repository struct{}
)

func (r *AwsS3Repository) Upload(value string, path models.Path, awsS3 models.AwsS3Param) error {
	svc := r.createS3Sess(awsS3)

	filePath := path.GetFullFilePath()

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(awsS3.Bucket),
		Key:         aws.String(filePath),
		Body:        bytes.NewReader([]byte(value)),
		ContentType: aws.String("ContentType"),
		Metadata: map[string]*string{
			"Key": aws.String("MetadataValue"),
		},
	})

	if err != nil {
		return errors.Wrap(err, "s3.Upload")
	}

	return nil
}

func (r *AwsS3Repository) createS3Sess(awsS3 models.AwsS3Param) *s3.S3 {
	cre := r.createCredentials(awsS3)
	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(awsS3.Region),
	})
	return cli
}

func (r *AwsS3Repository) createCredentials(awsS3 models.AwsS3Param) (cre *credentials.Credentials) {
	cre = credentials.NewStaticCredentials(
		awsS3.AwsId,
		awsS3.AwsKey,
		"",
	)
	return cre
}

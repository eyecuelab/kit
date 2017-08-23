package s3

import (
	"net/http"

	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/eyecuelab/kit/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
)

var uploder *s3manager.Uploader

func UploadFromForm(fileHeader *multipart.FileHeader, key string) (*s3manager.UploadOutput, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	io.Copy(&buffer, file)

	return Upload(buffer.Bytes(), key)
}

func Upload(b []byte, key string) (*s3manager.UploadOutput, error) {

	fileBytes := bytes.NewReader(b)
	fileType := http.DetectContentType(b)

	params := &s3manager.UploadInput{
		Bucket:      aws.String(viper.GetString("aws_bucket")),
		Key:         aws.String(key),
		Body:        fileBytes,
		ContentType: aws.String(fileType),
	}

	return uploder.Upload(params)
}

func newS3Client() *s3.S3 {
	awsAccessKey := viper.GetString("aws_access_key")
	awsSecret := viper.GetString("aws_secret")
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, "")
	_, err := creds.Get()
	log.Check(err)

	session, err := session.NewSession()
	log.Check(err)

	cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)
	return s3.New(session, cfg)
}

func setUploader() {
	uploder = s3manager.NewUploaderWithClient(newS3Client())
}

func init() {
	cobra.OnInitialize(setUploader)
}

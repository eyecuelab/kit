package s3

import (
	"net/http"
	"time"

	"bytes"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/eyecuelab/kit/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Bucket:      aws.String(viper.GetString("aws_bucket_name")),
		Key:         aws.String(key),
		Body:        fileBytes,
		ContentType: aws.String(fileType),
	}

	return uploder.Upload(params)
}

// Presign presign s3 key fora given period of time
func Presign(key string, duration time.Duration) (string, error) {
	svc := newS3Client()
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(viper.GetString("aws_bucket_name")),
		Key:    &key,
	})
	return req.Presign(duration)
}

func newS3Client() *s3.S3 {
	awsAccessKey := viper.GetString("aws_access_key")
	awsSecret := viper.GetString("aws_secret")
	region := viper.GetString("aws_bucket_location")

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, "")
	_, err := creds.Get()
	log.Check(err)

	session, err := session.NewSession()
	log.Check(err)

	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	return s3.New(session, cfg)
}

func setUploader() {
	uploder = s3manager.NewUploaderWithClient(newS3Client())
}

func init() {
	cobra.OnInitialize(setUploader)
}

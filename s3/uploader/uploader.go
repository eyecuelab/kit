package uploader

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// Config s3 uploader config data type
type Config struct {
	ID             string `jsonapi:"primary,uploader_config"`
	Key            string `jsonapi:"attr,key"`
	ACL            string `jsonapi:"attr,acl"`
	AwsAccessKeyID string `jsonapi:"attr,awsAccessKeyId"`
	S3Policy       string `jsonapi:"attr,s3Policy"`
	S3Signature    string `jsonapi:"attr,s3Signature"`
	S3Bucket       string `jsonapi:"attr,s3Bucket"`
	// ContentType    string `jsonapi:"attr,contentType"`
}

// Policy data type
type Policy struct {
	Expiration string        `json:"expiration"`
	Conditions []interface{} `json:"conditions"`
}

// ACL ...
func ACL() string {
	return "public-read"
}

// PolicyString ...
func PolicyString(key string) string {
	policyData := GetPolicy(key)
	policyDataJSON, _ := json.Marshal(policyData)
	return base64.StdEncoding.EncodeToString(policyDataJSON)
}

// Signature policy signature
func Signature(policy string) string {
	sigKey := []byte(viper.GetString("aws_secret"))

	sig := hmac.New(sha1.New, sigKey)
	sig.Write([]byte(policy))
	return base64.StdEncoding.EncodeToString(sig.Sum(nil))
}

// Key s3 file key
func Key(pref string) string {
	key, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s/", pref, key)
}

// GetPolicy generate policy for the file
func GetPolicy(key string) Policy {
	data := Policy{
		time.Now().UTC().Add(time.Hour * 10).Format(time.RFC3339),
		[]interface{}{
			[]string{"starts-with", "$key", key},
			map[string]string{"bucket": viper.GetString("aws_bucket_name")},
			map[string]string{"acl": ACL()},
			map[string]string{"success_action_status": "201"},
			[]interface{}{"content-length-range", 1, 104857600}, // TODO: use max size based on the type of the resource
		},
	}
	// TODO: validate content type based on the resource
	// conditions << [ 'starts-with', '$Content-Type', contentType]

	return data
}

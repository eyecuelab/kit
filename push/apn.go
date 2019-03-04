package push

import (
	"fmt"
	"net/http"

	"github.com/eyecuelab/kit/config"
	"github.com/eyecuelab/kit/log"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	alertTemplate = `{"aps":{"alert":"%v"}}`
	defautFile    = `/etc/eyecue_keys/push.p12`
)

var (
	client *apns2.Client
	topic  string
)

func init() {
	cobra.OnInitialize(setup, connectClient)
}

func setup() {
	viper.SetDefault("push_key_file", defautFile)
}

func connectClient() {
	topic = config.RequiredString("push_topic")
	keyFile := viper.GetString("push_key_file")

	if keyFile == "" {
		log.Warnf("No push key location provided")
		return
	}

	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		log.Warnf("No push key found at: %s", keyFile)
		return
	}

	cert, err := certificate.FromP12File(keyFile, "")

	if err != nil {
		log.Fatal(err)
	}
	client = apns2.NewClient(cert).Production()
}

func push(token string, payload string) error {
	notification := &apns2.Notification{
		DeviceToken: token,
		Topic:       topic,
		Payload:     payload,
	}

	res, err := client.Push(notification)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("push error: %v", res.Reason)
	}
	return nil
}

func Alert(token string, message string) error {
	return push(token, fmt.Sprintf(alertTemplate, message))
}

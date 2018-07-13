package push

import (
	"fmt"
	"net/http"

	"github.com/eyecuelab/kit/config"
	"github.com/eyecuelab/kit/log"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/spf13/cobra"
)

const (
	alertTemplate = `{"aps":{"alert":"%v"}}`
)

var (
	client *apns2.Client
	topic  string
)

func init() {
	cobra.OnInitialize(connectClient)
}

func connectClient() {
	topic = config.RequiredString("push_topic")

	cert, err := certificate.FromP12File("./lub_dev_push.p12", "")

	if err != nil {
		log.Fatal(err)
	}
	client = apns2.NewClient(cert).Development()
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

package sendgrid

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eyecuelab/kit/mailman"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

type jsonError map[string][]map[string]interface{}

type SGConfig struct {
	*mailman.Config
	ApiKey     string
	TemplateId string
}

func (c *SGConfig) FromEmail() *mail.Email {
	return mail.NewEmail(c.Config.From.Name, c.Config.From.Email)
}

type Mailer struct {
	*SGConfig
	SGClient *sendgrid.Client
}

func (m *Mailer) Config() *mailman.Config {
	return m.SGConfig.Config
}

func (m *Mailer) Send(to *mailman.Address, content *mailman.Content, vars *mailman.MergeVars) error {
	f := m.SGConfig.FromEmail()
	t := mail.NewEmail(to.Name, to.Email)

	html := mail.NewContent("text/html", content.HtmlBody.String())
	text := mail.NewContent("text/plain", content.PlainBody.String())

	message := mail.NewV3MailInit(f, content.Subject.String(), t, text)
	message.AddContent(html)
	message.SetTemplateID(m.SGConfig.TemplateId)

	for key, value := range vars.BodyVars {
		message.Personalizations[0].SetSubstitution(fmt.Sprintf("-%s-", key), value)
	}

	if response, err := m.SGClient.Send(message); err != nil {
		return err
	} else if response.StatusCode > 299 {
		return errors.New(parseError(response))
	}

	return nil
}

func MinConfig(fromName string, fromEmail string, domain string) *SGConfig {
	from := &mailman.Address{fromName, fromEmail}
	return &SGConfig{
		Config: &mailman.Config{from, domain},
	}
}

func Configure(config *SGConfig) {
	if len(config.ApiKey) == 0 {
		config.ApiKey = viper.GetString("sendgrid_api_key")
	}
	if len(config.TemplateId) == 0 {
		config.TemplateId = viper.GetString("sendgrid_template")
	}
	client := sendgrid.NewSendClient(config.ApiKey)
	mailman.Configure(&Mailer{config, client})
}

func parseError(response *rest.Response) string {
	var raw jsonError

	if err := json.Unmarshal([]byte(response.Body), &raw); err != nil {
		return response.Body
	}
	var msg string
	if len(raw["errors"]) > 0 {
		msg = raw["errors"][0]["message"].(string)
	} else {
		msg = response.Body
	}
	return fmt.Sprintf("SendGrid Error: %d, %s", response.StatusCode, msg)
}

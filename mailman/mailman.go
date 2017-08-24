package mailman

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/eyecuelab/kit/assets"
)

type Address struct {
	Name  string
	Email string
}

type Config struct {
	From   *Address
	Domain string
}

type Mailer interface {
	Config() *Config
	Send(to *Address, content *Content, vars *MergeVars) error
}

type Template struct {
	Subject   *template.Template
	HtmlBody  *template.Template
	PlainBody *template.Template
}

type Content struct {
	Subject   *bytes.Buffer
	HtmlBody  *bytes.Buffer
	PlainBody *bytes.Buffer
}

type MergeVars struct {
	SubjectVars map[string]string
	BodyVars    map[string]string
}

type mergeVars interface{}

var (
	TemplateDir string
	Templates   map[string]*Template
	mailer      Mailer
)

func newTemplate(s string) (*template.Template, error) {
	return template.New("").Parse(s)
}

func setDefaultTemplateDir() {
	dir := path.Join("data", "mail")

	if _, err := assets.Dir(dir); err == nil {
		TemplateDir = dir
	}
}

func Configure(m Mailer) {
	setDefaultTemplateDir()

	Templates = make(map[string]*Template)
	mailer = m
}

func AddTemplate(key string, subject string, htmlBody string, plainBody string) error {
	var s, h, p *template.Template
	var err error
	if s, err = newTemplate(subject); err != nil {
		return err
	}
	if h, err = newTemplate(htmlBody); err != nil {
		return err
	}
	if p, err = newTemplate(plainBody); err != nil {
		return err
	}

	Templates[key] = &Template{s, h, p}
	return nil
}

func readEmailTemplates(key string) ([][]byte, error) {
	templates := make([][]byte, 3)
	var err error

	for i, part := range []string{"subject", "html", "plain"} {
		fileName := fmt.Sprintf("%s_%s.tmpl", key, part)
		p := path.Join(TemplateDir, fileName)

		if templates[i], err = ioutil.ReadFile(p); err != nil {
			return templates, err
		}
	}

	return templates, nil
}

func RegisterTemplate(key string) error {
	if len(TemplateDir) == 0 {
		return errors.New("TemplateDir not set")
	}

	parts, err := readEmailTemplates(key)
	if err != nil {
		return err
	}

	return AddTemplate(key, string(parts[0]), string(parts[1]), string(parts[2]))
}

func Send(to *Address, templateKey string, vars *MergeVars) error {
	if mailer == nil {
		return errors.New("Mail has not been configured")
	}

	template := Templates[templateKey]
	if template == nil {
		return errors.New(fmt.Sprintf("Template not found for key [%s]", templateKey))
	}

	content := &Content{
		new(bytes.Buffer),
		new(bytes.Buffer),
		new(bytes.Buffer),
	}

	if err := template.Subject.Execute(content.Subject, vars.SubjectVars); err != nil {
		return err
	}
	if err := template.HtmlBody.Execute(content.HtmlBody, vars.BodyVars); err != nil {
		return err
	}
	if err := template.PlainBody.Execute(content.PlainBody, vars.BodyVars); err != nil {
		return err
	}

	return mailer.Send(to, content, vars)
}

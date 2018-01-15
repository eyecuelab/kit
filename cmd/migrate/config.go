package migrate

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"text/template"

	"github.com/eyecuelab/kit/goenv"
	"github.com/spf13/viper"
)

const configPath = "./dbconfig.yml"

type configData struct {
	Env         string
	DatabaseUrl string
}

func writeConfig() error {
	f, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EEXIST {
			return nil
		}
		return fmt.Errorf("writeConfig: os.OpenFile: %v", err)
	}
	defer f.Close()
	if err := templateToFile(f); err != nil {
		return fmt.Errorf("writeConfig: templateToFile: %v", err)
	}
	return nil
}

func templateToFile(wr io.Writer) error {
	tmpl, err := template.New("config").Parse(configTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(wr, getConfig())
}

func getConfig() configData {
	return configData{goenv.Env, viper.GetString("database_url")}
}

var configTemplate = `
{{- .Env }}:
    dialect: postgres
    datasource: {{ .DatabaseUrl }}
    dir: data/bin/migrations
`

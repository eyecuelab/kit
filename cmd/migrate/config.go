package migrate

import (
	"io"
	"os"
	"syscall"
	"text/template"

	"github.com/eyecuelab/kit/goenv"
	"github.com/eyecuelab/kit/log"
	"github.com/spf13/viper"
)

const configPath = "./dbconfig.yml"

func writeConfig() {
	f, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EEXIST {
			return
		}
		log.Fatalf("writeConfig: os.OpenFile: %v", err)
	}
	defer f.Close()

	templateToFile(f)
}

func templateToFile(wr io.Writer) {
	tmpl, err := template.New("config").Parse(configTemplate)
	log.Check(err)

	err = tmpl.Execute(wr, getConfig())
	log.Check(err)
}

func getConfig() configData {
	return configData{goenv.Env, viper.GetString("database_url")}
}

type configData struct {
	Env         string
	DatabaseUrl string
}

var configTemplate = `
{{- .Env }}:
    dialect: postgres
    datasource: {{ .DatabaseUrl }}
    dir: data/bin/migrations
`

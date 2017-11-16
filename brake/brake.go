package brake

import (
	"net/http"

	"github.com/airbrake/gobrake"
	"github.com/eyecuelab/kit/goenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Airbrake *gobrake.Notifier
	Env string
)

const traceDepth = 5

func init() {
	cobra.OnInitialize(setup)
}

func setup() {
	key := viper.GetString("airbrake_key")
	project := viper.GetInt64("airbrake_project")
	Env := viper.GetString("airbrake_env")

	if len(key) != 0 && project != 0 {
		Airbrake = gobrake.NewNotifier(project, key)
	}
}

func IsSetup() bool {
	return Airbrake != nil
}

func Notify(e error, req *http.Request) {
	if IsSetup() {
		notice := gobrake.NewNotice(e, req, traceDepth)
		notice.Context["environment"] = Env
		Airbrake.SendNotice(notice)
	}
}
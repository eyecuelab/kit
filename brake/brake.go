package brake

import (
	"net/http"

	"github.com/airbrake/gobrake"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Airbrake *gobrake.Notifier
)

func init() {
	cobra.OnInitialize(setup)
}

func setup() {
	key := viper.GetString("airbrake_key")
	project := viper.GetInt64("airbrake_project")

	if len(key) != 0 && project != 0 {
		Airbrake = gobrake.NewNotifier(project, key)
	}
}

func IsSetup() bool {
	return Airbrake != nil
}

func Notify(e error, req http.Request) {
	if IsSetup() {
		Airbrake.Notify(err, req)
	}
}

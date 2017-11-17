package brake

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/airbrake/gobrake"
	"github.com/eyecuelab/kit/functools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Airbrake *gobrake.Notifier
	Env      string
)

type severity string

const (
	traceDepth = 5

	SeverityError    severity = "error"
	SeverityWarn     severity = "warning"
	SeverityCritical severity = "critical"
)

func init() {
	cobra.OnInitialize(setup)
}

func setup() {
	key := viper.GetString("airbrake_key")
	project := viper.GetInt64("airbrake_project")
	Env = viper.GetString("airbrake_env")

	if len(key) != 0 && project != 0 {
		Airbrake = gobrake.NewNotifier(project, key)
	}
}

func IsSetup() bool {
	return Airbrake != nil
}

func Notify(e error, req *http.Request, sev severity) {
	if IsSetup() {
		notice := gobrake.NewNotice(e, req, traceDepth)
		setNoticeVars(notice, req, sev)
		Airbrake.SendNotice(notice)
	}
}

func setNoticeVars(n *gobrake.Notice, req *http.Request, sev severity) {
	n.Context["environment"] = Env
	n.Context["severity"] = sev
	if expectBody(req) {
		n.Params["body"] = body(req)
	}
}

func body(req *http.Request) interface{} {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return fmt.Sprintf("error reading body: %v", err)
	}

	if len(b) == 0 {
		return "no body"
	}

	if !isJsonContentType(req) {
		return string(b)
	}

	// if !json.Valid(b) {
	// 	return fmt.Sprintf("body is not valid JSON: %s", b)
	// }

	formatted := make(map[string]interface{})
	json.Unmarshal(b, &formatted)

	return formatted
}

func isJsonContentType(req *http.Request) bool {
	cType := req.Header.Get("Content-Type")
	cType = strings.ToLower(cType)
	return strings.Contains(cType, "json")
}

func expectBody(req *http.Request) bool {
	requestsWithBody := []string{http.MethodPost, http.MethodPatch, http.MethodPut}
	return functools.StringSliceContains(requestsWithBody, req.Method)
}

package facebook

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
)

const (
	Fields        = "first_name, last_name, email, timezone, gender, picture.height(800)"
	GraphEndpoint = "https://graph.facebook.com/me"
)

type GraphInfo struct {
	Error     graphError
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string
	Picture   picture
	Email     string
	Timezone  int
	Gender    string
}

type graphError struct {
	Message   string
	Type      string
	Code      int
	FbTraceId string `json:"fbtrace_id"`
}

type picture struct {
	Data struct {
		Url string
	}
}

func FbUserInfo(accessToken string) (GraphInfo, error) {
	var gInfo GraphInfo

	request := gorequest.New().Get(GraphEndpoint)

	_, body, errs := request.Param("access_token", accessToken).Param("fields", Fields).End()

	if errs != nil {
		return gInfo, errs[0]
	}

	err := json.Unmarshal([]byte(body), &gInfo)
	if gInfo.Error.Message != "" {
		return gInfo, errors.New(gInfo.Error.Message)
	}
	return gInfo, err
}

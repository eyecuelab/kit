package facebook

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

const (
	Fields        = "first_name, last_name, email, timezone, gender, picture"
	GraphEndpoint = "https://graph.facebook.com/me"
)

type GraphInfo struct {
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string
	Picture   picture
	Email     string
	Timezone  int
	Gender    string
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
	return gInfo, err
}

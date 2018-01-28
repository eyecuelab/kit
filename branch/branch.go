package branch

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/eyecuelab/kit/config"
	"github.com/parnurzeal/gorequest"
)

const branchApiUrl = "https://api.branch.io/v1/url"

type (
	request struct {
		BranchKey string `json:"branch_key"`
		Data      Data   `json:"data"`
	}

	Data struct {
		CanonicalUrl string `json:"$canonical_url"`
	}

	CanonicalUrlData struct {
		Url    string
		Params map[string]string
	}
	Response struct {
		Url   string
		Error branchError
	}

	branchError struct {
		Code    int
		Message string
	}
)

func GetUrl(cud CanonicalUrlData) (Response, error) {
	cURL, err := generateFullUrl(cud)
	if err != nil {
		return Response{}, err
	}

	return send(request{Data: Data{cURL}})
}

func generateFullUrl(cud CanonicalUrlData) (string, error) {
	baseURL, err := url.Parse(cud.Url)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	for k, v := range cud.Params {
		params.Add(k, v)
	}

	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}

func send(req request) (resp Response, err error) {
	req.BranchKey = config.RequiredString("BRANCH_KEY")

	request := gorequest.New()
	_, body, _ := request.Post(branchApiUrl).Send(req).End()

	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return resp, err
	}

	if resp.Error.Code > 299 {
		return resp, fmt.Errorf("Branch.io error: %v - %v", resp.Error.Code, resp.Error.Message)
	}
	return resp, nil
}

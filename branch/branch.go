package branch

import (
	"encoding/json"
	"fmt"
	"github.com/eyecuelab/kit/config"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

const branchApiUrl = "https://api.branch.io/v1/url"

type branchRequest struct {
	BranchKey string `json:"branch_key"`
	Data      Data   `json:"data"`
}

type Data struct {
	CanonicalUrl string `json:"$canonical_url"`
}

type CanonicalUrlData struct {
	Url    string
	Params map[string]string
}

type BranchResponse struct {
	Url   string
	Error branchError
}

type branchError struct {
	Code    int
	Message string
}

func GetUrl(cud CanonicalUrlData) (BranchResponse, error) {
	cUrl, err := generateFullUrl(cud)
	if err != nil {
		return BranchResponse{}, err
	}

	req := branchRequest{
		Data: Data{cUrl},
	}

	return send(req)
}

func generateFullUrl(cud CanonicalUrlData) (string, error) {
	baseUrl, err := url.Parse(cud.Url)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	for k, v := range cud.Params {
		params.Add(k, v)
	}

	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}

func send(req branchRequest) (BranchResponse, error) {
	req.BranchKey = config.RequiredString("BRANCH_KEY")

	request := gorequest.New()
	_, body, _ := request.Post(branchApiUrl).Send(req).End()

	var bResp BranchResponse
	if err := json.Unmarshal([]byte(body), &bResp); err != nil {
		return bResp, err
	}

	if bResp.Error.Code > 299 {
		return bResp, fmt.Errorf("Branch.io error: %v - %v", bResp.Error.Code, bResp.Error.Message)
	}
	return bResp, nil
}

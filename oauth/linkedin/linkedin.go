package linkedin

import (
	"encoding/json"
	"errors"

	"github.com/parnurzeal/gorequest"
)

const (
	// DataEndpoint linkedin profile data endpoint url
	DataEndpoint = "https://api.linkedin.com/v1/people/~:(id,first-name,last-name,email-address,picture-url,location,positions,headline)"
)

var tokenTypes = map[string]string{
	"oauth":  "oauth_token",
	"oauth2": "oauth2_access_token",
}

// User linkedin user data
type User struct {
	ID         string
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	PictureURL string `json:"pictureUrl"`
	Headline   string
	Positions  UserPositions
	Location   struct {
		Country struct {
			Code string
		}
		Name string
	}
	ErrorCode int `json:"errorCode"`
	Message   string
	Status    int
}

// UserPositions linkedin user postions data
type UserPositions struct {
	Values []UserPosition
}

// UserPosition linkedin user position data
type UserPosition struct {
	ID        int
	Title     string
	IsCurrent bool `json:"isCurrent"`
	Company   Company
	Location  struct {
		Country struct {
			Code string
			Name string
		}
	}
	StartDate struct {
		Month int
		Year  int
	} `json:"startDate"`
}

// Company linkedin company data
type Company struct {
	ID       int
	Name     string
	Industry string
	Size     string
	Type     string
}

// UserInfo linkedin user info by the access token
func UserInfo(tokenType string, accessToken string) (*User, error) {
	user := new(User)

	request := gorequest.New().Get(DataEndpoint)
	_, body, errs := request.
		Param(tokenTypes[tokenType], accessToken).
		Param("format", "json").
		End()
	if errs != nil {
		return user, errs[0]
	}

	if err := json.Unmarshal([]byte(body), &user); err != nil {
		return nil, err
	}

	if user.Message != "" {
		return user, errors.New(user.Message)
	}

	return user, nil
}

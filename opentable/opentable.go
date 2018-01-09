package opentable

import (
	"os"

	"github.com/parnurzeal/gorequest"
)

var _ = gorequest.GET

const tokenKey = "OPENTABLE_ACCESS_TOKEN"
const endpoint = "https://opentable.herokuapp.com/api"

//BadRequestError represents invalid request parameter(s).
type BadRequestError string

func (err BadRequestError) Error() string { return string(err) }

const (
	errBadPage      BadRequestError = "page must be >= 0"
	errBadPrice     BadRequestError = "price must be blank or in the closed interval [0, 4]"
	errBadPerPage   BadRequestError = "entries per page must be in one of the following: {5, 10, 15, 25, 50, 100}"
	errEmptyRequest BadRequestError = "request must specify at least one parameter"
)

func accessToken() (string, bool) {
	return os.LookupEnv(tokenKey)
}

type Query struct {
	Offset  int    `json:"offset,omitempty" bson:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty" bson:"limit,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
}

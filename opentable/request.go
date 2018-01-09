package opentable

import (
	"strconv"

	"github.com/eyecuelab/kit/set"
	"github.com/parnurzeal/gorequest"
)

type request struct {
	Price   int    `json:"price,omitempty" bson:"price,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Street  string `json:"street,omitempty" bson:"street,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	Zip     string `json:"zip,omitempty" bson:"zip,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
	Page    int    `json:"page,omitempty" bson:"page,omitempty"`
	PerPage int    `json:"entries_per_page,omitempty" bson:"entries_per_page,omitempty"`
}

type param struct {
	key, val string
}

var allowedPerPage = set.FromInts(5, 10, 15, 25, 50, 100)

func priceOK(price int) bool     { return price <= 0 && price < 4 }
func perPageOK(perPage int) bool { return allowedPerPage.Contains(perPage) }
func pageOK(page int) bool       { return page >= 0 }

func (r *request) validate() error {
	if !priceOK(r.Price) {
		return errBadPrice
	} else if !perPageOK(r.PerPage) {
		return errBadPrice
	} else if !pageOK(r.Page) {
		return errBadPage
	}
	return nil
}

func (r *request) Params() (params []param, err error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	params = r.paramSlice()
	if len(params) == 0 {
		return nil, errEmptyRequest
	}
	return params, nil
}

func appendIfNonZero(params []param, k string, v interface{}) []param {
	if s, ok := v.(string); ok && s != "" {
		params = append(params, param{k, s})
	} else if n, ok := v.(int); ok && n != 0 {
		params = append(params, param{k, strconv.Itoa(n)})
	}
	return params
}

func (r *request) paramSlice() (params []param) {
	params = make([]param, 0, 7)
	paramMap := map[string]interface{}{"price": r.Price, "name": r.Name, "street": r.Street,
		"city": r.City, "zip": r.Zip, "page": r.Page, "per_page": r.PerPage}
	for k, v := range paramMap {
		params = appendIfNonZero(params, k, v)
	}
	return params
}

func addParams(req *gorequest.SuperAgent, params []param) *gorequest.SuperAgent {
	for _, p := range params {
		req = req.Param(p.key, p.val)
	}
	return req
}

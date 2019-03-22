package pagination

import (
	"math"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/url"

)

type (
	Pagination struct {
		Count   int `json:"item_count"`
		Max     int `json:"max"`
		Page    int `json:"page"`
		PerPage int
		Offset  int
		Url   url.URL
	}

)

func (p *Pagination) Links() map[string]string {
	pageValues := map[string]int{ "self": p.Page }

	if p.Max != 1 {
		pageValues["first"] = 1
		pageValues["last"] = p.Max
	}

	if p.Max != p.Page {
		pageValues["next"] = p.Page + 1
	}

	if p.Page > 1 {
		pageValues["prev"] = p.Page - 1
	}

	return p.linkify(pageValues)
}

func (p Pagination) linkify(pageValues map[string]int) map[string]string {
	links := make(map[string]string)
	for k, v := range pageValues {
		links[k] = p.withPageNumber(v)
	}
	return links
}

func (p Pagination) withPageNumber(num int) string {
	q := p.Url.Query()
	q.Set("page[number]", strconv.Itoa(num))
	p.Url.RawQuery = q.Encode()

	return p.Url.RequestURI()
}

// DefaultPerPage ...
var DefaultPerPage = 20

// MaxPerPage ...
var MaxPerPage = 100

// Data pagination data for jsonapi
var Data Pagination

// Apply apply pagination to a provided scope
func Apply(c echo.Context, scope *gorm.DB, model interface{}, list interface{}, perPage int) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}

	var count int
	if err := scope.Offset(0).Model(model).Count(&count).Error; err != nil {
		return err
	}

	strPerPage := c.QueryParam("per_page")
	if strPerPage != "" {
		perPage, _ = strconv.Atoi(strPerPage)
	} else if perPage == 0 {
		perPage = DefaultPerPage
	}

	if perPage > MaxPerPage {
		perPage = MaxPerPage
	}

	setData(count, perPage, page)

	if err := scope.Offset(offset(page, perPage)).Limit(perPage).
		Find(list).Error; err != nil {
		return err
	}

	return nil
}

func setData(count int, perPage int, page int) {
	Data.Count = count
	pages := float64(count) / float64(perPage)
	Data.Max = int(math.Ceil(pages))
	Data.Page = page
}

func offset(pageNumber int, perPage int) int {
	return (pageNumber - 1) * perPage
}

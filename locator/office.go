package locator

import (
	"net/url"
)

type Office struct {
	Slug string
	Name string
	Lat  float64
	Long float64
}

func (o Office) URL(baseURL url.URL) url.URL {
	path, _ := url.Parse("/" + o.Slug)
	return *baseURL.ResolveReference(path)
}

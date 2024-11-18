package main

import (
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/vinta/pangu"
	"guthub.com/baaj2109/go-house/model"
)

const (
	rootURL = "https://rent.591.com.tw/"
)

// GenerateURL is convert options to query parameters.
func GenerateURL(o *Options) (string, error) {
	v, error := query.Values(o)

	return rootURL + "?" + v.Encode(), error
}

func stringReplacer(text string) string {
	replacer := strings.NewReplacer("\n", "", " ", "")

	return pangu.SpacingText(replacer.Replace(text))
}

func trimTextSpace(s string) string {
	return strings.Fields(s)[0]
}

func fillDescription(s []string) []string {
	s = append(s, s[2])
	s[2] = "沒有格局說明"

	return s
}

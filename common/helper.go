package common

import (
	"regexp"
	"strings"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func MakeSlug(text string) string {

	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}

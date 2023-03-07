package common

import (
	"database/sql/driver"
	"regexp"
	"strings"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type GenderType string

const (
	MEN   GenderType = "men"
	WOMEN GenderType = "women"
)

func (gt *GenderType) Scan(value interface{}) error {
	*gt = GenderType(value.([]byte))
	return nil
}

func (gt GenderType) Value() (driver.Value, error) {
	return string(gt), nil
}

func MakeSlug(text string) string {

	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}

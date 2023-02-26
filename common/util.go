package common

import (
	"strconv"
	"strings"
)

func ContainsString(items []string, ele string) bool {
	if items == nil {
		return false
	}

	for _, item := range items {
		if item == ele {
			return true
		}
	}

	return false
}

func ContainsInt(items []int, ele int) bool {
	if items == nil {
		return false
	}
	for _, item := range items {
		if item == ele {
			return true
		}
	}
	return false
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ConvertToInt(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return 0, err
	} else {
		return num, nil
	}
}

func ConvertToBool(str string) (bool, error) {
	num, err := strconv.ParseBool(str)

	if err != nil {
		return false, err
	} else {
		return num, nil
	}
}

func ConvertQueryToStrings(str string) (results []string) {
	for _, s := range strings.Split(str, ",") {
		results = append(results, s)
	}
	return
}

type StallInfo interface {
	GetMarketCode() string
	GetFloorCode() string
	GetStallCode() string
}

func NewBool(val bool) *bool {
	b := val
	return &b
}

type CheckOrNumber struct {
	OrNumber    string `json:"or_number"`
	ReferenceId string `json:"reference_id"`
}

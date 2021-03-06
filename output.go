package exec

import (
	"strconv"
	"strings"
)

type output struct {
	text string
	err  error
}

func (o output) HasError() bool {
	if o.err != nil {
		return true
	}

	return false
}

func (o output) String() string {
	return o.text
}

func (o output) Int() int {
	i, err := strconv.Atoi(o.text)
	if err == nil {
		return i
	}
	return 0
}

func (o output) Bool() bool {
	if "true" == o.text {
		return true
	}

	return false
}

func (o output) Slice(sep string) []string {
	return strings.Split(o.text, sep)
}

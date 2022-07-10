package lib

import (
	"bytes"
	"encoding/json"
)

type Formatter interface {
	JSON(input string) string
	YAML(input string) string
	XML(input string) string
}

type Format struct {
}

func (f *Format) JSON(input string) string {
	var out bytes.Buffer
	json.Indent(&out, []byte(input), "", "  ")
	return string(out.Bytes())
}

func (f *Format) YAML(input string) string {
	panic("not implemented") // TODO: Implement
}

func (f *Format) XML(input string) string {
	panic("not implemented") // TODO: Implement
}

func NewFormat() Formatter {
	return &Format{}
}

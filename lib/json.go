package lib

import (
	"encoding/json"

	"github.com/tidwall/pretty"
)

type JSONFormater interface {
	TabIndent(string) string
	TwoSpacesIndent(string) string
	FourSpacesIndent(string) string
	Minify(string) string
	IsJSON(string) bool
}

var (
	tab        = &pretty.Options{Indent: "	", Width: 80, Prefix: "", SortKeys: false}
	twoSpaces  = &pretty.Options{Indent: "  ", Width: 80, Prefix: "", SortKeys: false}
	fourSpaces = &pretty.Options{Indent: "    ", Width: 80, Prefix: "", SortKeys: false}
)

type JSONFormat struct {
}

func (j *JSONFormat) TabIndent(s string) string {
	return string(pretty.PrettyOptions([]byte(s), tab))
}

func (j *JSONFormat) TwoSpacesIndent(s string) string {
	return string(pretty.PrettyOptions([]byte(s), twoSpaces))
}

func (j *JSONFormat) FourSpacesIndent(s string) string {
	return string(pretty.PrettyOptions([]byte(s), fourSpaces))
}

func (j *JSONFormat) Minify(s string) string {
	return string(pretty.Ugly([]byte(s)))
}

func (j *JSONFormat) IsJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func NewJSONFormat() JSONFormater {
	return &JSONFormat{}
}

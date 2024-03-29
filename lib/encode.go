package lib

import (
	b64 "encoding/base64"
	"html"
	"net/url"
)

type Encoder interface {
	Base64(input string) string
	URL(input string) string
	HTML(input string) string
}

type Encode struct {
}

func (e *Encode) Base64(input string) string {
	return b64.StdEncoding.EncodeToString([]byte(input))
}

func (e *Encode) URL(input string) string {
	return url.PathEscape(input)
}

func (e *Encode) HTML(input string) string {
	return html.EscapeString(input)
}

func NewEncoder() Encoder {
	return &Encode{}
}

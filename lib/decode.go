package lib

import (
	b64 "encoding/base64"
	"html"
	"log"
	"net/url"
)

type Decoder interface {
	Base64(input string) string
	URL(input string) string
	HTML(input string) string
}

type Decode struct {
}

func (e *Decode) Base64(input string) string {
	res, err := b64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(res)
}

func (e *Decode) URL(input string) string {
	res, _ := url.QueryUnescape(input)
	return res
}

func (e *Decode) HTML(input string) string {
	return html.UnescapeString(input)
}

func NewDecoder() Decoder {
	return &Decode{}
}

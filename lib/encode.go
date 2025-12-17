package lib

import (
	b64 "encoding/base64"
	"fmt"
	"html"
	"net/url"
	"strings"
)

type Encoder interface {
	Base64(input string) string
	URL(input string) string
	HTML(input string) string
	Unicode(input string) string
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

func (e *Encode) Unicode(input string) string {
	var b strings.Builder
	for _, r := range input {
		// ASCII range is left untouched.
		if r <= 0x7F {
			b.WriteRune(r)
			continue
		}

		// BMP characters.
		if r <= 0xFFFF {
			fmt.Fprintf(&b, "\\u%04X", r)
			continue
		}

		// Characters outside BMP are encoded as UTF-16 surrogate pairs.
		r1 := 0xD800 + ((r-0x10000)>>10)&0x3FF
		r2 := 0xDC00 + (r-0x10000)&0x3FF
		fmt.Fprintf(&b, "\\u%04X\\u%04X", r1, r2)
	}
	return b.String()
}

func NewEncoder() Encoder {
	return &Encode{}
}

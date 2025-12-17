package lib

import (
	b64 "encoding/base64"
	"encoding/hex"
	"html"
	"log"
	"net/url"
	"strings"
	"unicode/utf16"
)

type Decoder interface {
	Base64(input string) string
	URL(input string) string
	HTML(input string) string
	Unicode(input string) string
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

func (e *Decode) Unicode(input string) string {
	var b strings.Builder

	for i := 0; i < len(input); {
		if i+6 <= len(input) && input[i] == '\\' && input[i+1] == 'u' {
			hexDigits := input[i+2 : i+6]
			buf, err := hex.DecodeString(hexDigits)
			if err != nil || len(buf) != 2 {
				// Not valid hex, write the literal and move on.
				b.WriteByte(input[i])
				i++
				continue
			}

			codeUnit := int(buf[0])<<8 | int(buf[1])

			// Handle surrogate pairs.
			if 0xD800 <= codeUnit && codeUnit <= 0xDBFF && i+12 <= len(input) && input[i+6] == '\\' && input[i+7] == 'u' {
				hexDigits2 := input[i+8 : i+12]
				buf2, err2 := hex.DecodeString(hexDigits2)
				if err2 == nil && len(buf2) == 2 {
					codeUnit2 := int(buf2[0])<<8 | int(buf2[1])
					if 0xDC00 <= codeUnit2 && codeUnit2 <= 0xDFFF {
						r := utf16.DecodeRune(rune(codeUnit), rune(codeUnit2))
						b.WriteRune(r)
						i += 12
						continue
					}
				}
			}

			// Single code unit.
			b.WriteRune(rune(codeUnit))
			i += 6
			continue
		}

		b.WriteByte(input[i])
		i++
	}

	return b.String()
}

func NewDecoder() Decoder {
	return &Decode{}
}

package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func MD5(input string) string {
	m := md5.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func SHA1(input string) string {
	m := sha1.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func SHA256(input string) string {
	m := sha256.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func SHA512(input string) string {
	m := sha512.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

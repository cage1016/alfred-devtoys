package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type Hasher interface {
	MD5(input string) string
	SHA1(input string) string
	SHA256(input string) string
	SHA512(input string) string
}

type Hash struct {
}

func (h *Hash) MD5(input string) string {
	m := md5.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func (h *Hash) SHA1(input string) string {
	m := sha1.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func (h *Hash) SHA256(input string) string {
	m := sha256.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func (h *Hash) SHA512(input string) string {
	m := sha512.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func NewHasher() Hasher {
	return &Hash{}
}

package lib

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

type CheckSumer interface {
	MD5() string
	SHA1() string
	SHA256() string
	SHA512() string
}

type CheckSum struct {
	Md5    string
	Sha1   string
	Sha256 string
	Sha512 string
}

func (cm *CheckSum) MD5() string {
	return cm.Md5
}

func (cm *CheckSum) SHA1() string {
	return cm.Sha1
}

func (cm *CheckSum) SHA256() string {
	return cm.Sha256
}

func (cm *CheckSum) SHA512() string {
	return cm.Sha512
}

func NewCheckSum(file string) (CheckSumer, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	md5 := md5.New()
	sha1 := sha1.New()
	sha256 := sha256.New()
	sha512 := sha512.New()

	// For optimum speed, Getpagesize returns the underlying system's memory page size.
	pagesize := os.Getpagesize()

	// wraps the Reader object into a new buffered reader to read the files in chunks
	// and buffering them for performance.
	reader := bufio.NewReaderSize(f, pagesize)

	// creates a multiplexer Writer object that will duplicate all write
	// operations when copying data from source into all different hashing algorithms
	// at the same time
	multiWriter := io.MultiWriter(md5, sha1, sha256, sha512)

	// Using a buffered reader, this will write to the writer multiplexer
	// so we only traverse through the file once, and can calculate all hashes
	// in a single byte buffered scan pass.
	//
	_, err = io.Copy(multiWriter, reader)
	if err != nil {
		return nil, err
	}

	return &CheckSum{
		Md5:    hex.EncodeToString(md5.Sum(nil)),
		Sha1:   hex.EncodeToString(sha1.Sum(nil)),
		Sha256: hex.EncodeToString(sha256.Sum(nil)),
		Sha512: hex.EncodeToString(sha512.Sum(nil)),
	}, nil
}

package lib

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/gabriel-vasile/mimetype"
)

const MaxImageSize = 2 << 20

var ErrFileTooLarge = errors.New("file is too large")

func ImageEncode(image string) (string, string, error) {
	fi, err := statRegular(image)
	if err != nil {
		return "", "", err
	}
	if fi.Size() > MaxImageSize {
		return "", "", ErrFileTooLarge
	}

	bytes, err := ioutil.ReadFile(image)
	if err != nil {
		return "", "", err
	}

	mtype := mimetype.Detect(bytes)
	return base64.StdEncoding.EncodeToString(bytes), mtype.String(), nil
}

func Download(url, dataDir string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %s", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch image: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(io.LimitReader(resp.Body, MaxImageSize+1))
	if err != nil {
		return "", fmt.Errorf("failed to read image: %s", err)
	}
	if len(bodyBytes) > MaxImageSize {
		return "", ErrFileTooLarge
	}

	_, format, err := image.DecodeConfig(bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/d.%s", dataDir, format)
	file, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, (bytes.NewReader(bodyBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %s", err)
	}

	return path, nil
}

package lib

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/gabriel-vasile/mimetype"
)

func ImageEncode(image string) (string, string, error) {
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
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch image: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}

	buff := bytes.NewBuffer(nil)
	bodyBytes, err := ioutil.ReadAll(io.TeeReader(resp.Body, buff))
	if err != nil {
		return "", fmt.Errorf("failed to read image: %s", err)
	}

	_, format, err := image.DecodeConfig(buff)
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

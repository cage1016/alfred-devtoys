/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/alfred"
)

// qrcodeCmd represents the qrcode command
var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "QR codes generator",
	Run:   runQrcode,
}

func createQRCodeByBoombuler(content string, quality qr.ErrorCorrectionLevel, size int, dest string) (err error) {
	qrCode, err := qr.Encode(content, quality, qr.Auto)
	if err != nil {
		return
	}

	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return
	}

	file, err := os.Create(dest)
	if err != nil {
		return
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		return
	}

	return
}

func runQrcode(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	log.Println(query)

	CheckForUpdate()

	path := fmt.Sprintf("%s/qr.png", wf.DataDir())
	s, err := strconv.Atoi(alfred.GetQrcodeSize(wf))
	if err != nil {
		wf.NewItem(err.Error()).Subtitle("QR Code Config Size fail").Valid(false).Icon(aw.IconError)
	} else {
		if err != createQRCodeByBoombuler(query, qr.M, s, path) {
			wf.NewItem(err.Error()).Subtitle("QR Code").Valid(false).Icon(aw.IconError)
		} else {
			wf.NewItem(query).
				Subtitle("⌘+L ⇧, ↩ Browse QR code file").
				Valid(true).
				Quicklook(path).
				Largetype(query).
				Arg(path).
				Icon(&aw.Icon{Value: path}).
				Var("action", "browse in alfred")
		}
	}
	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(qrcodeCmd)
}

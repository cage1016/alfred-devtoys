/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/lib"
)

var reUrl = regexp.MustCompile(`(?m)^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)

// imageB64Cmd represents the imageB64 command
var imageB64Cmd = &cobra.Command{
	Use:   "imageB64",
	Short: "Image Base64 Encode",
	Run:   runImageB64,
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func runImageB64(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	var err error
	if reUrl.Match([]byte(query)) {
		query, err = lib.Download(query, wf.DataDir())
		if err != nil {
			wf.NewItem(fmt.Sprintf("`%s` download fail", query)).Subtitle("Try a different query?").Icon(Base64ImgGrayIcon)
			wf.SendFeedback()
			return
		}
	}

	b64EncodeStr, mtype, err := lib.ImageEncode(query)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid file", query)).Subtitle("Try a different query?").Icon(Base64ImgGrayIcon)
	} else {
		wf.NewItem(b64EncodeStr).
			Subtitle("⌘+L ⇧, ↩ Copy Base64 string").
			Valid(true).
			Quicklook(query).
			Largetype(b64EncodeStr).
			Icon(Base64ImgIcon).
			Arg(b64EncodeStr).
			Var("action", "copy")

		dataURI := fmt.Sprintf("data:%s;base64,%s", mtype, b64EncodeStr)
		wf.NewItem(dataURI).
			Subtitle("⌘+L ⇧, ↩ Copy Base64 Data URI").
			Valid(true).
			Quicklook(query).
			Largetype(dataURI).
			Icon(Base64ImgIcon).
			Arg(dataURI).
			Var("action", "copy")

		str2 := fmt.Sprintf("<img src=\"%s\">", dataURI)
		wf.NewItem(str2).
			Subtitle("⌘+L ⇧, ↩ Copy HTML <img> code").
			Valid(true).
			Quicklook(query).
			Largetype(str2).
			Icon(Base64ImgIcon).
			Arg(str2).
			Var("action", "copy")

		str3 := fmt.Sprintf("background-image: url(\"%s\");", dataURI)
		wf.NewItem(str3).
			Subtitle("⌘+L ⇧, ↩ Copy CSS Background Source").
			Valid(true).
			Quicklook(query).
			Largetype(str3).
			Icon(Base64ImgIcon).
			Arg(str3).
			Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(imageB64Cmd)
}

/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"strings"

	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/lib"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode string",
	Run:   encodeRun,
}

func encodeRun(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	coder := lib.NewEncoder()
	b64EncodeStr := coder.Base64(query)
	wf.NewItem(b64EncodeStr).
		Subtitle("⌘+L, ↩ Copy Base64 Encode").
		Valid(true).
		Largetype(b64EncodeStr).
		Icon(Base64Icon).
		Arg(b64EncodeStr).
		Var("action", "copy")

	urlEncodeStr := coder.URL(query)
	wf.NewItem(urlEncodeStr).
		Subtitle("⌘+L, ↩ Copy URL Encode").
		Valid(true).
		Largetype(urlEncodeStr).
		Icon(UrlIcon).
		Arg(urlEncodeStr).
		Var("action", "copy")

	htmlEncodeStr := coder.HTML(query)
	wf.NewItem(htmlEncodeStr).
		Subtitle("⌘+L, ↩ Copy HTML Encode").
		Valid(true).
		Arg(htmlEncodeStr).
		Largetype(htmlEncodeStr).
		Icon(HtmlIcon).
		Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}

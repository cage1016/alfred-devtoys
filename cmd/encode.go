/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode string",
	Run:   encodeRun,
}

func encodeRun(cmd *cobra.Command, args []string) {
	coder := lib.NewEncoder()
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	b64EncodeStr := coder.Base64(query)
	wf.NewItem(b64EncodeStr).Subtitle("Base64 Encode").Valid(true).Icon(Base64Icon).Arg(b64EncodeStr).Var("action", "copy")

	urlEncodeStr := coder.URL(query)
	wf.NewItem(urlEncodeStr).Subtitle("URL Encode").Valid(true).Icon(UrlIcon).Arg(urlEncodeStr).Var("action", "copy")

	htmlEncodeStr := coder.HTML(query)
	wf.NewItem(htmlEncodeStr).Subtitle("HTML Encode").Valid(true).Arg(htmlEncodeStr).Icon(HtmlIcon).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode string",
	Run:   decodeRun,
}

func decodeRun(cmd *cobra.Command, args []string) {
	coder := lib.NewDecoder()
	str := strings.Join(args, " ")
	if strings.TrimSpace(str) == "" {
		str = string(clipboard.Read(clipboard.FmtText))
	}

	b64DecodeStr := coder.Base64(str)
	wf.NewItem(b64DecodeStr).Subtitle("Base64 Decode").Valid(true).Arg(b64DecodeStr).Icon(Base64Icon).Var("action", "copy")

	urlDecodeStr := coder.URL(str)
	wf.NewItem(urlDecodeStr).Subtitle("URL Decode").Valid(true).Arg(urlDecodeStr).Icon(UrlIcon).Var("action", "copy")

	htmlDecodeStr := coder.HTML(str)
	wf.NewItem(htmlDecodeStr).Subtitle("HTML Decode").Valid(true).Arg(htmlDecodeStr).Icon(HtmlIcon).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

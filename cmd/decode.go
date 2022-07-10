/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strings"

	"github.com/cage1016/alfred-devtoys/lib"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
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
	wf.NewItem(b64DecodeStr).Subtitle("Base64 Decode").Valid(true).Arg(b64DecodeStr).Var("action", "copy")

	urlDecodeStr := coder.URL(str)
	wf.NewItem(urlDecodeStr).Subtitle("URL Decode").Valid(true).Arg(urlDecodeStr).Var("action", "copy")

	htmlDecodeStr := coder.HTML(str)
	wf.NewItem(htmlDecodeStr).Subtitle("HTML Decode").Valid(true).Arg(htmlDecodeStr).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

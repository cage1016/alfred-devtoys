/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode string",
	Run:   encodeRun,
}

func encodeRun(cmd *cobra.Command, args []string) {
	coder := lib.NewEncoder()
	str := strings.Join(args, " ")
	if strings.TrimSpace(str) == "" {
		str = string(clipboard.Read(clipboard.FmtText))
	}

	b64EncodeStr := coder.Base64(str)
	wf.NewItem(b64EncodeStr).Subtitle("Base64 Encode").Valid(true).Icon(&aw.Icon{Value: "base64.pdf"}).Arg(b64EncodeStr).Var("action", "copy")

	urlEncodeStr := coder.URL(str)
	wf.NewItem(urlEncodeStr).Subtitle("URL Encode").Valid(true).Icon(&aw.Icon{Value: "url.pdf"}).Arg(urlEncodeStr).Var("action", "copy")

	htmlEncodeStr := coder.HTML(str)
	wf.NewItem(htmlEncodeStr).Subtitle("HTML Encode").Valid(true).Arg(htmlEncodeStr).Icon(&aw.Icon{Value: "html.pdf"}).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}

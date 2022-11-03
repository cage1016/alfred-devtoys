/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/lib"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode string",
	Run:   decodeRun,
}

func decodeRun(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	log.Println(query)

	CheckForUpdate()

	coder := lib.NewDecoder()
	b64DecodeStr := coder.Base64(query)
	wf.NewItem(b64DecodeStr).
		Subtitle("⌘+L, ↩ Copy Base64 Decode").
		Valid(true).
		Arg(b64DecodeStr).
		Largetype(b64DecodeStr).
		Icon(Base64Icon).
		Var("action", "copy")

	urlDecodeStr := coder.URL(query)
	wf.NewItem(urlDecodeStr).
		Subtitle("⌘+L, ↩ Copy URL Decode").
		Valid(true).
		Arg(urlDecodeStr).
		Largetype(urlDecodeStr).
		Icon(UrlIcon).
		Var("action", "copy")

	htmlDecodeStr := coder.HTML(query)
	wf.NewItem(htmlDecodeStr).
		Subtitle("⌘+L, ↩ Copy HTML Decode").
		Valid(true).
		Arg(htmlDecodeStr).
		Largetype(htmlDecodeStr).
		Icon(HtmlIcon).
		Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

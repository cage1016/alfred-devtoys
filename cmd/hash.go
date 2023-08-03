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

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash string",
	Run:   runHash,
}

func runHash(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	md5Str := lib.MD5(query)
	wf.NewItem(md5Str).
		Subtitle("⌘+L, ↩ Copy MD5").
		Valid(true).
		Arg(md5Str).
		Largetype(md5Str).
		Icon(HashIcon).
		Var("action", "copy")

	sha1Str := lib.SHA1(query)
	wf.NewItem(sha1Str).
		Subtitle("⌘+L, ↩ Copy SHA1").
		Valid(true).
		Arg(sha1Str).
		Largetype(sha1Str).
		Icon(HashIcon).
		Var("action", "copy")

	sha256Str := lib.SHA256(query)
	wf.NewItem(sha256Str).
		Subtitle("⌘+L, ↩ Copy SHA256").
		Valid(true).
		Arg(sha256Str).
		Largetype(sha256Str).
		Icon(HashIcon).
		Var("action", "copy")

	sha512Str := lib.SHA512(query)
	wf.NewItem(sha512Str).
		Subtitle("⌘+L, ↩ Copy SHA512").
		Valid(true).
		Arg(sha512Str).
		Largetype(sha512Str).
		Icon(HashIcon).
		Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

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

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "hash string",
	Run:   runHash,
}

func runHash(cmd *cobra.Command, args []string) {
	coder := lib.NewHasher()
	str := strings.Join(args, " ")
	if strings.TrimSpace(str) == "" {
		str = string(clipboard.Read(clipboard.FmtText))
	}

	md5Str := coder.MD5(str)
	wf.NewItem(md5Str).Subtitle("MD5").Valid(true).Arg(md5Str).Icon(HashIcon).Var("action", "copy")

	sha1Str := coder.SHA1(str)
	wf.NewItem(sha1Str).Subtitle("SHA1").Valid(true).Arg(sha1Str).Icon(HashIcon).Var("action", "copy")

	sha256Str := coder.SHA256(str)
	wf.NewItem(sha256Str).Subtitle("SHA256").Valid(true).Arg(sha256Str).Icon(HashIcon).Var("action", "copy")

	sha512Str := coder.SHA512(str)
	wf.NewItem(sha512Str).Subtitle("SHA512").Valid(true).Arg(sha512Str).Icon(HashIcon).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

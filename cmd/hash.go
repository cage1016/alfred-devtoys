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

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash string",
	Run:   runHash,
}

func runHash(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	md5Str := lib.MD5(query)
	wf.NewItem(md5Str).Subtitle("MD5").Valid(true).Arg(md5Str).Icon(HashIcon).Var("action", "copy")

	sha1Str := lib.SHA1(query)
	wf.NewItem(sha1Str).Subtitle("SHA1").Valid(true).Arg(sha1Str).Icon(HashIcon).Var("action", "copy")

	sha256Str := lib.SHA256(query)
	wf.NewItem(sha256Str).Subtitle("SHA256").Valid(true).Arg(sha256Str).Icon(HashIcon).Var("action", "copy")

	sha512Str := lib.SHA512(query)
	wf.NewItem(sha512Str).Subtitle("SHA512").Valid(true).Arg(sha512Str).Icon(HashIcon).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

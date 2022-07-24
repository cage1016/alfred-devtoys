/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
)

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "File checksum",
	Run:   runChecksum,
}

func runChecksum(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	s, err := lib.NewCheckSum(query)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid file", query)).Subtitle("Try a different query?").Icon(HashGrayIcon)
	} else {
		wf.NewItem(s.MD5()).Subtitle("MD5 CheckSum").Valid(true).Icon(HashIcon).Arg(s.MD5()).Var("action", "copy")
		wf.NewItem(s.SHA1()).Subtitle("SHA1 CheckSum").Valid(true).Icon(HashIcon).Arg(s.SHA1()).Var("action", "copy")
		wf.NewItem(s.SHA256()).Subtitle("SHA256 CheckSum").Valid(true).Icon(HashIcon).Arg(s.SHA256()).Var("action", "copy")
		wf.NewItem(s.SHA512()).Subtitle("SHA512 CheckSum").Valid(true).Icon(HashIcon).Arg(s.SHA512()).Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(checksumCmd)
}

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

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "file checksum",
	Run:   runChecksum,
}

func runChecksum(cmd *cobra.Command, args []string) {
	str := strings.Join(args, " ")
	if strings.TrimSpace(str) == "" {
		str = string(clipboard.Read(clipboard.FmtText))
	}

	s, err := lib.NewCheckSum(str)
	if err != nil {
		wf.NewItem(err.Error()).Subtitle("File CheckSum Fail").Valid(false).Icon(aw.IconError)
	} else {
		wf.NewItem(s.MD5()).Subtitle("MD5 CheckSum").Valid(true).Icon(&aw.Icon{Value: "hash.pdf"}).Arg(s.MD5()).Var("action", "copy")
		wf.NewItem(s.SHA1()).Subtitle("SHA1 CheckSum").Valid(true).Icon(&aw.Icon{Value: "hash.pdf"}).Arg(s.SHA1()).Var("action", "copy")
		wf.NewItem(s.SHA256()).Subtitle("SHA256 CheckSum").Valid(true).Icon(&aw.Icon{Value: "hash.pdf"}).Arg(s.SHA256()).Var("action", "copy")
		wf.NewItem(s.SHA512()).Subtitle("SHA512 CheckSum").Valid(true).Icon(&aw.Icon{Value: "hash.pdf"}).Arg(s.SHA512()).Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(checksumCmd)
}

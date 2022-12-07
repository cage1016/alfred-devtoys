/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/lib"
)

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "File checksum",
	Run:   runChecksum,
}

func runChecksum(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	CheckForUpdate()

	s, err := lib.NewCheckSum(query)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid file", query)).Subtitle("Try a different query?").Icon(HashGrayIcon)
	} else {
		wf.NewItem(s.MD5()).
			Subtitle("⌘+L ⇧, ↩ Copy MD5 CheckSum").
			Valid(true).
			Arg(s.MD5()).
			Quicklook(query).
			Largetype(s.MD5()).
			Icon(HashIcon).
			Var("action", "copy")

		wf.NewItem(s.SHA1()).
			Subtitle("⌘+L ⇧, ↩ Copy SHA1 CheckSum").
			Valid(true).
			Arg(s.SHA1()).
			Quicklook(query).
			Largetype(s.SHA1()).
			Icon(HashIcon).
			Var("action", "copy")

		wf.NewItem(s.SHA256()).
			Subtitle("⌘+L ⇧, ↩ Copy SHA256 CheckSum").
			Valid(true).
			Arg(s.SHA256()).
			Quicklook(query).
			Largetype(s.SHA256()).
			Icon(HashIcon).
			Var("action", "copy")

		wf.NewItem(s.SHA512()).
			Subtitle("⌘+L ⇧, ↩ Copy SHA512 CheckSum").
			Valid(true).
			Arg(s.SHA512()).
			Quicklook(query).
			Largetype(s.SHA512()).
			Icon(HashIcon).
			Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(checksumCmd)
}

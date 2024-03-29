/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/alfred"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "UUID string",
	Run:   runUuid,
}

func runUuid(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
		_, err := strconv.ParseUint(query, 10, 64)
		if err != nil {
			query = alfred.GetUuidDefault(wf)
		}
	}
	logrus.Debugf("query: %s", query)

	c, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid integer", query)).Subtitle("Try a different query?").Icon(UuidGrayIcon)
	} else {
		buf := make([]string, c)
		for i := int64(0); i < c; i++ {
			buf[i] = uuid.New().String()
		}
		for _, v := range buf {
			wf.NewItem(v).
				Subtitle("⌘+L ⌥, ↩ Copy UUID").
				Valid(true).
				Largetype(v).
				Arg(v).
				Icon(UuidIcon).
				Var("action", "copy").
				Valid(true).
				Opt().
				Subtitle("↩ Copy all of UUIDs").
				Arg(strings.Join(buf, "\n")).
				Var("action", "copy")
		}
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}

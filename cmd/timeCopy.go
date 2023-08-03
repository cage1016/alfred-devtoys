/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/cage1016/alfred-devtoys/alfred"
	"github.com/spf13/cobra"
)

// timeCopyCmd represents the timeCopy command
var timeCopyCmd = &cobra.Command{
	Use:   "timeCopy",
	Short: "Unix time convert with more time format",
	Run:   runTimeCopyCmd,
}

func runTimeCopyCmd(cmd *cobra.Command, args []string) {

	prefFile := path.Join(wf.Dir(), "info.plist")
	if dpf := os.Getenv("DEBUG_PLIST_FOLDER"); dpf != "" {
		prefFile = path.Join(dpf, "info.plist")
	}
	preferences := alfred.LoadPlist(prefFile)

	tfis := []TimeFormatItem{}
	userconfigurationconfig := preferences["userconfigurationconfig"].([]interface{})
	for _, v := range userconfigurationconfig {
		cfg := v.(map[string]interface{})
		if cfg["variable"].(string) == alfred.TimeFormat {
			for _, pairs := range cfg["config"].(map[string]interface{})["pairs"].([]interface{}) {
				pair := pairs.([]interface{})
				tfis = append(tfis, TimeFormatItem{
					Name:  pair[0].(string),
					Value: pair[1].(string),
				})
			}
		}
	}

	timeZone, _ := cmd.Flags().GetString("timeZone")
	tl, _ := time.LoadLocation(timeZone)

	if query, err := strconv.ParseInt(args[0], 10, 64); err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid", args[0])).Subtitle("Try a different query?").Icon(TimeConverterGrayIcon)
	} else {
		tm := time.Unix(query, 0).UTC()
		for _, t := range tfis {
			wf.NewItem(tm.In(tl).Format(t.Name)).
				Subtitle(fmt.Sprintf("⌘+L, ↩ Copy Local (%s), Format: %s", timeZone, t.Value)).
				Arg(t.Value).
				Largetype(tm.In(tl).Format(t.Name)).
				Icon(TimeConverterIcon).
				Valid(true)
		}
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(timeCopyCmd)
	timeCopyCmd.PersistentFlags().StringP("timeZone", "t", "", "time zone")
}

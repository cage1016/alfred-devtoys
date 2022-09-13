/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tj/go-naturaldate"

	"github.com/cage1016/alfred-devtoys/alfred"
	"github.com/cage1016/alfred-devtoys/lib"
)

const defaultTimeFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Unix time converter",
	Run:   runTimeCmd,
}

func runTimeCmd(cmd *cobra.Command, args []string) {
	// log.Println(query)
	CheckForUpdate()

	var query int64
	var err error
	var tm time.Time

	if query, err = strconv.ParseInt(args[0], 10, 64); err == nil {
		tm = time.Unix(query, 0).UTC()
	} else {
		tm, err = naturaldate.Parse(strings.ToLower(args[0]), time.Now(), naturaldate.WithDirection(naturaldate.Future))
	}

	if err != nil {
		wf.NewItem(fmt.Sprintf("`%d` is invalid", query)).Subtitle("Try a different query?").Icon(TimeConverterGrayIcon)
	} else {
		timeFormat := alfred.GetTimeFormat(wf)
		if timeFormat == "" {
			timeFormat = defaultTimeFormat
		}

		timeZone := alfred.GetTimeZone(wf)
		if timeZone != "" {
			tl, err := time.LoadLocation(timeZone)
			if err != nil {
				wf.NewItem(fmt.Sprintf("Invalid time zone: %s", timeZone)).Subtitle("Try a different time zone?").Icon(TimeConverterGrayIcon)
			} else {
				l := tm.In(tl).Format(timeFormat)
				wi := wf.NewItem(l).
					Subtitle(fmt.Sprintf("⌘+L ⌥ ^, ↩ Copy Local (%s), Format: %s", timeZone, timeFormat)).
					Arg(l).
					Valid(true).
					Largetype(l).
					Icon(TimeConverterIcon).
					Var("action", "copy")

				wi.Ctrl().
					Subtitle("↩ Visit go-naturaldate to check more natural date example").
					Arg("https://github.com/tj/go-naturaldate/blob/v1.3.0/naturaldate_test.go").
					Valid(true).
					Var("action", "browse")

				wi.Opt().
					Subtitle("↩ Visit https://golang.org/pkg/time/#pkg-constants to check more time format").
					Arg("https://golang.org/pkg/time/#pkg-constants").
					Valid(true).
					Var("action", "browse")
			}
		}

		{
			u := tm.UTC().Format(timeFormat)
			wi := wf.NewItem(u).
				Subtitle(fmt.Sprintf("⌘+L ⌥ ^, ↩ Copy UTC (ISO 8601), Format: %s", timeFormat)).
				Arg(u).
				Valid(true).
				Largetype(u).
				Icon(TimeConverterIcon).
				Var("action", "copy")

			wi.Ctrl().
				Subtitle("↩ Visit go-naturaldate to check more natural date example").
				Valid(true).
				Arg("https://github.com/tj/go-naturaldate/blob/v1.3.0/naturaldate_test.go").
				Var("action", "browse")

			wi.Opt().
				Subtitle("↩ Visit https://golang.org/pkg/time/#pkg-constants to check more time format").
				Valid(true).
				Arg("https://golang.org/pkg/time/#pkg-constants").
				Var("action", "browse")
		}

		u2 := fmt.Sprintf("%d", tm.Unix())
		wf.NewItem(u2).
			Subtitle("⌘+L, ↩ Copy Unix Timestamp").
			Arg(u2).
			Valid(true).
			Largetype(u2).
			Icon(TimeConverterIcon)

		d := lib.TimeDuration(time.Since(tm))
		wf.NewItem(d).
			Valid(true).
			Largetype(d).
			Icon(TimeConverterIcon)
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(timeCmd)
}

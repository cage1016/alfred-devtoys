/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-devtoys/lib"
)

// caseCmd represents the case command
var caseCmd = &cobra.Command{
	Use:   "case",
	Short: "Change string case",
	Run:   runCase,
}

func runCase(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	cases := lib.NewChangeCase()
	t, _ := cmd.Flags().GetString("type")
	if t == "command" {
		val := reflect.ValueOf(*cases)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name
			if len(query) == 0 {
				wf.NewItem(fmt.Sprintf("`%s` is invalid input", query)).Subtitle(fmt.Sprintf("Try a different query for %s?", name)).Icon(TextChangeCaseGrayIcon)
			} else {
				cc := val.Field(i).Interface().(lib.ChangeCaser)
				str := cc.Fn(query)
				wf.NewItem(str).
					Subtitle(fmt.Sprintf("%s ➜ ⌘+L, ↩ Copy %s", name, cc.Subtitle())).
					Valid(true).
					Arg(str).
					Largetype(str).
					Icon(TextChangeCaseIcon).
					Var("action", "copy")
			}
		}
	} else {
		if cc := cases.Get(t); cc != nil {
			if len(query) == 0 {
				wf.NewItem(fmt.Sprintf("`%s` is invalid input", query)).Subtitle(fmt.Sprintf("Try a different query for %s?", t)).Icon(TextChangeCaseGrayIcon)
			} else {
				str := cc.Fn(query)
				wf.NewItem(str).
					Subtitle(fmt.Sprintf("%s ➜ ⌘+L, ↩ Copy %s", t, cc.Subtitle())).
					Valid(true).
					Arg(str).
					Largetype(str).
					Icon(TextChangeCaseIcon).
					Var("action", "copy")
			}
		}
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(caseCmd)

	caseCmd.PersistentFlags().StringP("type", "t", "", "type of case")
}

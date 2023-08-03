/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json",
	Run:   runJSON,
}

func runJSON(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	j := lib.NewJSONFormat()
	if j.IsJSON(query) {
		wf.NewItem(query).
			Subtitle("⌘+L, ↩ Action in Indent as Tab").
			Valid(true).
			Arg(j.TabIndent(query)).
			Largetype(j.TabIndent(query)).
			Icon(JsonIcon).
			Var("action", "action in alfred")

		wf.NewItem(query).
			Subtitle("⌘+L, ↩ Action in Indent as 2 Spaces").
			Valid(true).
			Arg(j.TwoSpacesIndent(query)).
			Largetype(j.TwoSpacesIndent(query)).
			Icon(JsonIcon).
			Var("action", "action in alfred")

		wf.NewItem(query).
			Subtitle("⌘+L, ↩ Action in Indent as 4 Spaces").
			Valid(true).
			Arg(j.FourSpacesIndent(query)).
			Largetype(j.FourSpacesIndent(query)).
			Icon(JsonIcon).
			Var("action", "action in alfred")

		wf.NewItem(query).
			Subtitle("⌘+L, ↩ Action in Minified").
			Valid(true).
			Arg(j.Minify(query)).
			Largetype(j.Minify(query)).
			Icon(JsonIcon).
			Var("action", "action in alfred")

	} else {
		wf.NewItem(fmt.Sprintf("`%s` is invalid JSON", query)).Subtitle("Try a different query?").Icon(JsonGrayIcon)
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}

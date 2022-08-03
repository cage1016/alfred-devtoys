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

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json",
	Run:   runJSON,
}

func runJSON(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	j := lib.NewJSONFormat()
	if j.IsJSON(query) {
		wf.NewItem(query).Subtitle("Indent as Tab").Valid(true).Arg(j.TabIndent(query)).Icon(JsonIcon).Var("action", "action in alfred")
		wf.NewItem(query).Subtitle("Indent as 2 Spaces").Valid(true).Arg(j.TwoSpacesIndent(query)).Icon(JsonIcon).Var("action", "action in alfred")
		wf.NewItem(query).Subtitle("Indent as 4 Spaces").Valid(true).Arg(j.FourSpacesIndent(query)).Icon(JsonIcon).Var("action", "action in alfred")
		wf.NewItem(query).Subtitle("Minified").Valid(true).Arg(j.Minify(query)).Icon(JsonIcon).Var("action", "action in alfred")
	} else {
		wf.NewItem(fmt.Sprintf("`%s` is invalid JSON", query)).Subtitle("Try a different query?").Icon(JsonGrayIcon)
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}
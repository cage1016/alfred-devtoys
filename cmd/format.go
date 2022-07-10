/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strings"

	"github.com/cage1016/alfred-devtoys/lib"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "format string",
	Run:   runFormat,
}

func runFormat(c *cobra.Command, args []string) {
	coder := lib.NewFormat()
	str := strings.Join(args, " ")
	if strings.TrimSpace(str) == "" {
		str = string(clipboard.Read(clipboard.FmtText))
	}

	j := coder.JSON(str)
	wf.NewItem(j).Subtitle("JSON").Valid(true).Arg(j).Var("action", "copy")

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(formatCmd)
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"strings"

	"github.com/cage1016/alfred-devtoys/lib"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "jwt decode",
	Run:   runJwt,
}

func runJwt(c *cobra.Command, args []string) {
	str := strings.Join(args, " ")
	if strings.TrimSpace(str) == "" {
		str = string(clipboard.Read(clipboard.FmtText))
	}

	token, err := lib.JWTdecode(str)
	if err == nil {
		a, _ := json.Marshal(token.Header)
		b, _ := json.Marshal(token.Claims)

		wf.NewItem(string(a)).Subtitle("Header").Valid(true).Arg(string(a)).Var("action", "copy")
		wf.NewItem(string(b)).Subtitle("Payload").Valid(true).Arg(string(b)).Var("action", "copy")
	} else {
		wf.WarnEmpty("Invalid token", err.Error())
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}

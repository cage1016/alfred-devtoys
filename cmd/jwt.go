/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
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

		wf.NewItem(string(a)).Subtitle("Header").Valid(true).Arg(string(a)).Icon(JwtIcon).Var("action", "copy")
		wf.NewItem(string(b)).Subtitle("Payload").Valid(true).Arg(string(b)).Icon(JwtIcon).Var("action", "copy")
	} else {
		wf.NewItem(err.Error()).Subtitle("JSON").Valid(false).Icon(aw.IconError)
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}

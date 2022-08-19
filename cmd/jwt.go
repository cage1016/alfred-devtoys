/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Jwt decode",
	Run:   runJwt,
}

func runJwt(c *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	token, err := lib.JWTdecode(query)
	if err == nil {
		a, _ := json.Marshal(token.Header)
		b, _ := json.Marshal(token.Claims)

		wf.NewItem(string(a)).
			Subtitle("⌘+L, ↩ Action in Header").
			Valid(true).
			Arg(string(a)).
			Largetype(string(a)).
			Icon(JwtIcon).
			Var("action", "action in alfred")

		wf.NewItem(string(b)).
			Subtitle("⌘+L, ↩ Action in Payload").
			Valid(true).
			Arg(string(b)).
			Largetype(string(b)).
			Icon(JwtIcon).
			Var("action", "action in alfred")
	} else {
		wf.NewItem(fmt.Sprintf("`%s` is invalid jwt", query)).Subtitle("Try a different query?").Icon(JwtGrayIcon)
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}

/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/loremipsum.v1"

	"github.com/cage1016/alfred-devtoys/alfred"
)

// loremCmd represents the lorem command
var loremCmd = &cobra.Command{
	Use:   "li",
	Short: "Lorem ipsum is a dummy text generator",
	Run:   runLorem,
}

func runLorem(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
		_, err := strconv.ParseUint(query, 10, 64)
		if err != nil {
			query = alfred.GetLiDefault(wf)
		}
	}
	logrus.Debugf("query: %s", query)

	c, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid integer", query)).Subtitle("Try a different query?").Icon(LoremIpsumGrayIcon)
	} else {
		loremIpsumGeneratoe := loremipsum.New()
		words := loremIpsumGeneratoe.Words(int(c))
		wf.NewItem(words).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %d Words", c)).
			Valid(true).
			Arg(words).
			Largetype(words).Icon(LoremIpsumIcon).
			Var("action", "copy").
			Valid(true)

		sentences := loremIpsumGeneratoe.Sentences(int(c))
		wf.NewItem(sentences).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %d Sentences", c)).
			Valid(true).
			Arg(sentences).
			Largetype(sentences).Icon(LoremIpsumIcon).
			Var("action", "copy")

		paragraphs := strings.Join(strings.Split(loremIpsumGeneratoe.Paragraphs(int(c)), `\n`), "\n\n")
		wf.NewItem(paragraphs).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %d Paragraphs", c)).
			Valid(true).
			Arg(paragraphs).
			Largetype(paragraphs).Icon(LoremIpsumIcon).
			Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(loremCmd)
}

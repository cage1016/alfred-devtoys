/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

const (
	loremMaxCount = 100
	loremAttempts = 5
)

func generateLorem(c int) (words, sentences, paragraphs string, ok bool) {
	for i := 0; i < loremAttempts && !ok; i++ {
		func() {
			defer func() { recover() }()
			g := loremipsum.NewWithSeed(time.Now().UnixNano())
			words = g.Words(c)
			sentences = g.Sentences(c)
			paragraphs = strings.Join(strings.Split(g.Paragraphs(c), `\n`), "\n\n")
			ok = true
		}()
	}
	return
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
		c = clamp(c, 1, loremMaxCount)
		words, sentences, paragraphs, ok := generateLorem(int(c))
		if !ok {
			wf.NewItem("Failed to generate lorem ipsum").Subtitle("Try again?").Icon(LoremIpsumGrayIcon)
			wf.SendFeedback()
			return
		}

		wf.NewItem(words).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %d Words", c)).
			Valid(true).
			Arg(words).
			Largetype(words).Icon(LoremIpsumIcon).
			Var("action", "copy").
			Valid(true)

		wf.NewItem(sentences).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %d Sentences", c)).
			Valid(true).
			Arg(sentences).
			Largetype(sentences).Icon(LoremIpsumIcon).
			Var("action", "copy")

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

/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
	"gopkg.in/loremipsum.v1"
)

// loremCmd represents the lorem command
var loremCmd = &cobra.Command{
	Use:   "li",
	Short: "Lorem ipsum is a dummy text generator",
	Run:   runLorem,
}

func runLorem(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	c, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid integer", query)).Subtitle("Try a different query?").Icon(LoremIpsumGrayIcon)
	} else {
		loremIpsumGeneratoe := loremipsum.New()
		words := loremIpsumGeneratoe.Words(int(c))
		wf.NewItem(words).Subtitle(fmt.Sprintf("%d Words", c)).Valid(true).Arg(words).Icon(LoremIpsumIcon).Var("action", "copy").Valid(true)

		sentences := loremIpsumGeneratoe.Sentences(int(c))
		wf.NewItem(sentences).Subtitle(fmt.Sprintf("%d Sentences", c)).Valid(true).Arg(sentences).Icon(LoremIpsumIcon).Var("action", "copy")

		paragraphs := strings.Join(strings.Split(loremIpsumGeneratoe.Paragraphs(int(c)), `\n`), "\n\n")
		wf.NewItem(paragraphs).Subtitle(fmt.Sprintf("%d Paragraphs", c)).Valid(true).Arg(paragraphs).Icon(LoremIpsumIcon).Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(loremCmd)
}

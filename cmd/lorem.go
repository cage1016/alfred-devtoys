/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
	"gopkg.in/loremipsum.v1"
)

// loremCmd represents the lorem command
var loremCmd = &cobra.Command{
	Use:   "li",
	Short: "lorem ipsum is a dummy text generator",
	Run:   runLorem,
}

func runLorem(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	c, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		wf.NewItem(err.Error()).Subtitle("Invalid integer").Valid(false).Icon(aw.IconError)
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

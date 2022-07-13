/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	changecase "github.com/ku/go-change-case"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var M = map[string]struct {
	Fn       func(string) string
	Subtitle string
}{
	"camel":    {changecase.Camel, "Convert to a string with the separators denoted by having the next letter capitalized"},
	"constant": {changecase.Constant, "Convert to an upper case, underscore separated string"},
	"dot":      {changecase.Dot, "Convert to a lower case, period separated string"},
	"lower":    {changecase.Lower, "Convert to a string in lower case"},
	"lcfirst":  {changecase.LcFirst, "Convert to a string with the first character lower cased"},
	"no":       {changecase.No, "Convert the string without any casing (lower case, space separated)"},
	"param":    {changecase.Param, "Convert to a lower case, dash separated string"},
	"pascal":   {changecase.Pascal, "Convert to a string denoted in the same fashion as camelCase, but with the first letter also capitalized"},
	"path":     {changecase.Path, "Convert to a lower case, slash separated string"},
	"sentence": {changecase.Sentence, "Convert to a lower case, space separated string"},
	"snake":    {changecase.Snake, "Convert to a lower case, underscore separated string"},
	"swap":     {changecase.Swap, "Convert to a string with every character case reversed"},
	"title":    {changecase.Title, "Convert to a space separated string with the first character of every word upper cased"},
	"upper":    {changecase.Upper, "Convert to a string in upper case"},
	"ucfirst":  {changecase.UcFirst, "Convert to a string with the first character upper cased"},
	"hashtag": {func(s string) string {
		array := regexp.MustCompile(" +").Split(s, -1)
		for i := 0; i < len(array); i++ {
			if len(array[i]) > 0 && array[i][0:1] != "#" {
				array[i] = "#" + strings.ToLower(array[i])
			}
		}
		return strings.Join(array, " ")
	}, "Convert to a string, space separated string with hashtag symbols"},
}

// caseCmd represents the case command
var caseCmd = &cobra.Command{
	Use:   "case",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		t, _ := cmd.Flags().GetString("type")

		query := strings.Join(args, " ")
		if strings.TrimSpace(query) == "" {
			query = string(clipboard.Read(clipboard.FmtText))
		}

		if m, ok := M[t]; ok {
			str := m.Fn(query)
			wf.NewItem(str).Subtitle(m.Subtitle).Valid(true).Arg(str).Icon(TextChangeCaseIcon)
		}

		if t == "list" {
			for k, v := range M {
				wf.NewItem(fmt.Sprintf("Change Case %s", changecase.UcFirst(k))).Subtitle(v.Subtitle).Valid(true).Icon(TextChangeCaseIcon).Var("action", k)
			}

			wf.Filter(args[0])
			wf.WarnEmpty("No matching items", "Try a different query?")
		}

		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(caseCmd)

	caseCmd.PersistentFlags().StringP("type", "t", "", "type of case")
}

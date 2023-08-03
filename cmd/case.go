/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/atotto/clipboard"
	changecase "github.com/ku/go-change-case"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	Short: "Change string case",
	Run:   runCase,
}

func runCase(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	t, _ := cmd.Flags().GetString("type")
	if t == "command" {
		keys := make([]string, 0, len(M))
		for k := range M {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, v := range keys {
			m := M[v]
			if len(query) == 0 {
				wf.NewItem(fmt.Sprintf("`%s` is invalid input", query)).Subtitle(fmt.Sprintf("Try a different query for %s?", v)).Icon(TextChangeCaseGrayIcon)
			} else {
				str := m.Fn(query)
				wf.NewItem(str).
					Subtitle(fmt.Sprintf("%s ➜ ⌘+L, ↩ Copy %s", v, m.Subtitle)).
					Valid(true).
					Arg(str).
					Largetype(str).
					Icon(TextChangeCaseIcon).
					Var("action", "copy")
			}
		}
	} else {
		if m, ok := M[t]; ok {
			if len(query) == 0 {
				wf.NewItem(fmt.Sprintf("`%s` is invalid input", query)).Subtitle(fmt.Sprintf("Try a different query for %s?", t)).Icon(TextChangeCaseGrayIcon)
			} else {
				str := m.Fn(query)
				wf.NewItem(str).
					Subtitle(fmt.Sprintf("%s ➜ ⌘+L, ↩ Copy %s", t, m.Subtitle)).
					Valid(true).
					Arg(str).
					Largetype(str).
					Icon(TextChangeCaseIcon).
					Var("action", "copy")
			}
		}
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(caseCmd)

	caseCmd.PersistentFlags().StringP("type", "t", "", "type of case")
}

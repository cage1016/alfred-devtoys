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

	"github.com/cage1016/alfred-devtoys/lib"
)

// octalCmd represents the octal command
var octalCmd = &cobra.Command{
	Use:   "octal",
	Short: "Octal To ➜",
	Run:   runOctal,
}

func runOctal(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	_, err := strconv.ParseUint(query, 8, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid octal", query)).Subtitle("Try a different query?").Icon(NumberGrayIcon)
	} else {
		b := lib.OctToBin(query)
		wf.NewItem(BinaryFormat(b)).
			Subtitle(" ⌘+L, ↩ Copy Octal to Binary").
			Valid(true).
			Arg(b).
			Largetype(b).
			Icon(NumberIcon).
			Var("action", "copy")

		o := lib.OctToDec(query)
		wf.NewItem(DecimalFormat(o)).
			Subtitle(" ⌘+L, ↩ Copy Octal to Decimal").
			Valid(true).
			Arg(o).
			Largetype(o).
			Icon(NumberIcon).
			Var("action", "copy")

		h := lib.OctToHex(query)
		wf.NewItem(HexFormat(h)).
			Subtitle(" ⌘+L, ↩ Copy Octal to Hexadecimal").
			Valid(true).
			Arg(h).
			Largetype(h).
			Icon(NumberIcon).
			Var("action", "copy")
	}
	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(octalCmd)
}

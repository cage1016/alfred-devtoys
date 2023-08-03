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

// binaryCmd represents the binary command
var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Binary To ➜",
	Run:   runBinary,
}

func runBinary(cmd *cobra.Command, args []string) {
	query := args[0]
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	_, err := strconv.ParseUint(query, 2, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid binary", query)).Subtitle("Try a different query?").Icon(NumberGrayIcon)
	} else {
		b := lib.BinToDec(query)
		wf.NewItem(DecimalFormat(b)).
			Subtitle(" ⌘+L, ↩ Copy Binary to Decimal").
			Valid(true).
			Arg(b).
			Largetype(b).
			Icon(NumberIcon).
			Var("action", "copy")

		o := lib.BinToOct(query)
		wf.NewItem(OctalFormat(o)).
			Subtitle(" ⌘+L, ↩ Copy Binary to Octal").
			Valid(true).
			Arg(o).
			Largetype(o).
			Icon(NumberIcon).
			Var("action", "copy")

		h := lib.BinToHex(query)
		wf.NewItem(HexFormat(h)).
			Subtitle(" ⌘+L, ↩ Copy Binary to Hexadecimal").
			Valid(true).
			Arg(h).
			Largetype(h).
			Icon(NumberIcon).
			Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(binaryCmd)
}

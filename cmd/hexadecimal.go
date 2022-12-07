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

// hexadecimalCmd represents the hexadecimal command
var hexadecimalCmd = &cobra.Command{
	Use:   "hexadecimal",
	Short: "Hexadecimal To ➜",
	Run:   runHexadecimal,
}

func runHexadecimal(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query, _ = clipboard.ReadAll()
	}
	logrus.Debugf("query: %s", query)

	CheckForUpdate()

	_, err := strconv.ParseUint(query, 16, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid hexadecimal", query)).Subtitle("Try a different query?").Icon(NumberGrayIcon)
	} else {
		b := lib.HexToBin(query)
		wf.NewItem(BinaryFormat(b)).
			Subtitle(" ⌘+L, ↩ Copy Hexadecimal to Binary").
			Valid(true).
			Arg(b).
			Largetype(b).
			Icon(NumberIcon).
			Var("action", "copy")

		o := lib.HexToOct(query)
		wf.NewItem(OctalFormat(o)).
			Subtitle(" ⌘+L, ↩ Copy Hexadecimal to Octal").
			Valid(true).
			Arg(o).
			Largetype(o).
			Icon(NumberIcon).
			Var("action", "copy")

		d := lib.HexToDec(query)
		wf.NewItem(DecimalFormat(d)).
			Subtitle(" ⌘+L, ↩ Copy Hexadecimal to Decimal").
			Valid(true).
			Arg(d).
			Largetype(d).
			Icon(NumberIcon).
			Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(hexadecimalCmd)
}

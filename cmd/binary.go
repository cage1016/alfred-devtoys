/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cage1016/alfred-devtoys/lib"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
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
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

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

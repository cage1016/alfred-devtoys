/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-devtoys/lib"
)

// decimalCmd represents the decimal command
var decimalCmd = &cobra.Command{
	Use:   "decimal",
	Short: "Decimal To ➜",
	Run:   runDecimal,
}

func runDecimal(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	c, err := strconv.ParseUint(query, 10, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid decimal", query)).Subtitle("Try a different query?").Icon(NumberGrayIcon)
	} else {
		if c > math.MaxInt64 {
			wf.NewItem(fmt.Sprintf("`%s` is overflows int", query)).Subtitle("Try a different query?").Icon(NumberGrayIcon)
		} else {
			b := lib.DecToBin(query)
			wf.NewItem(BinaryFormat(b)).
				Subtitle(" ⌘+L, ↩ Copy Decimal to Binary").
				Valid(true).
				Arg(b).
				Largetype(b).
				Icon(NumberIcon).
				Var("action", "copy")

			o := lib.DecToOct(query)
			wf.NewItem(OctalFormat(o)).
				Subtitle(" ⌘+L, ↩ Copy Decimal to Octal").
				Valid(true).
				Arg(o).
				Largetype(o).
				Icon(NumberIcon).
				Var("action", "copy")

			h := lib.DecToHex(query)
			wf.NewItem(HexFormat(h)).
				Subtitle(" ⌘+L, ↩ Copy Decimal to Hexadecimal").
				Valid(true).
				Arg(h).
				Largetype(h).
				Icon(NumberIcon).
				Var("action", "copy")
		}
	}
	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(decimalCmd)
}

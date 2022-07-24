/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	query := strings.Join(args, " ")
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
		wf.NewItem(DecimalFormat(b)).Subtitle("➜ Binary to Decimal").Valid(true).Arg(b).Icon(NumberIcon).Var("action", "copy")

		o := lib.BinToOct(query)
		wf.NewItem(OctalFormat(o)).Subtitle("➜ Binary to Octal").Valid(true).Arg(o).Icon(NumberIcon).Var("action", "copy")

		h := lib.BinToHex(query)
		wf.NewItem(HexFormat(h)).Subtitle("➜ Binary to Hexadecimal").Valid(true).Arg(h).Icon(NumberIcon).Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(binaryCmd)
}

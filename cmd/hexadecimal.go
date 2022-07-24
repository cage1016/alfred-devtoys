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

// hexadecimalCmd represents the hexadecimal command
var hexadecimalCmd = &cobra.Command{
	Use:   "hexadecimal",
	Short: "Hexadecimal To ➜",
	Run:   runHexadecimal,
}

func runHexadecimal(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}
	log.Println(query)

	CheckForUpdate()

	_, err := strconv.ParseUint(query, 16, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("`%s` is invalid hexadecimal", query)).Subtitle("Try a different query?").Icon(NumberGrayIcon)
	} else {
		b := lib.HexToBin(query)
		wf.NewItem(BinaryFormat(b)).Subtitle("➜ Hexadecimal to Binary").Valid(true).Arg(b).Icon(NumberIcon).Var("action", "copy")

		o := lib.HexToOct(query)
		wf.NewItem(OctalFormat(o)).Subtitle("➜ Hexadecimal to Octal").Valid(true).Arg(o).Icon(NumberIcon).Var("action", "copy")

		d := lib.HexToDec(query)
		wf.NewItem(DecimalFormat(d)).Subtitle("➜ Hexadecimal to Decimal").Valid(true).Arg(d).Icon(NumberIcon).Var("action", "copy")
	}

	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(hexadecimalCmd)
}

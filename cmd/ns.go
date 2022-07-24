/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

func format(in string, l int, sep rune) string {
	numOfDigits := len(in)
	numOfCommas := (numOfDigits - 1) / l
	out := make([]byte, len(in)+numOfCommas)

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == l {
			j, k = j-1, 0
			out[j] = byte(sep)
		}
	}
}

func DecimalFormat(in string) string {
	if len(in) > 3 {
		return format(in, 3, ',')
	}
	return in
}

func OctalFormat(in string) string {
	if len(in) > 3 {
		return format(in, 3, ' ')
	}
	return in
}

func BinaryFormat(in string) string {
	if len(in) > 4 {
		return format(in, 4, ' ')
	}
	return in
}

func HexFormat(in string) string {
	if len(in) > 4 {
		return format(in, 4, ' ')
	}
	return in
}

// nsCmd represents the nbc command
var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Number system convert",
}

func init() {
	rootCmd.AddCommand(nsCmd)
}

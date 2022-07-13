/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "uuid string",
	Run:   runUuid,
}

func runUuid(cmd *cobra.Command, args []string) {
	query := strings.Join(args, " ")
	c, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		wf.NewItem(err.Error()).Subtitle("Invalid integer").Valid(false).Icon(aw.IconError)
	} else {
		buf := make([]string, c)
		for i := int64(0); i < c; i++ {
			buf[i] = uuid.New().String()
		}
		for _, v := range buf {
			wf.NewItem(v).Subtitle("UUID, ⌥ to copy all").Valid(true).Arg(v).Icon(UuidIcon).Var("action", "copy").Valid(true).
				Opt().Subtitle(fmt.Sprintf("UUID, copy all of %d", c)).Arg(strings.Join(buf, "\n")).Var("action", "copy")
		}
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}

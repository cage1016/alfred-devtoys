/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// nscommandCmd represents the nscommand command
var nscommandCmd = &cobra.Command{
	Use:   "command",
	Short: "list number system commands",
	Run:   runNsCommand,
}

func runNsCommand(cmd *cobra.Command, args []string) {

	wf.NewItem("Binary To ➜").Subtitle("↩ Launch Decimal, Octal, Hexadecimal").Valid(true).UID("1").Icon(NumberIcon).Var("type", "binary")
	wf.NewItem("Decimal To ➜").Subtitle("↩ Launch Binary, Octal, Hexadecimal").Valid(true).UID("2").Icon(NumberIcon).Var("type", "decimal")
	wf.NewItem("Octal To ➜").Subtitle("↩ Launch Binary, Decimal, Hexadecimal").Valid(true).UID("3").Icon(NumberIcon).Var("type", "octal")
	wf.NewItem("Hexadecimal To ➜").Subtitle("↩ Launch Binary, Octal, Decimal").Valid(true).UID("4").Icon(NumberIcon).Var("type", "hexadecimal")

	wf.Filter(args[0])
	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

func init() {
	nsCmd.AddCommand(nscommandCmd)
}

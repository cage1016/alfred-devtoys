/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/spf13/cobra"
)

const updateJobName = "checkForUpdate"

var (
	repo = "cage1016/alfred-devtoys"
	wf   *aw.Workflow
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alfred-devtoys",
	Short: "A Swiss Army knife for Alfred",
	Run: func(cmd *cobra.Command, args []string) {

		if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
			log.Println("Running update check in background...")

			cmd := exec.Command(os.Args[0], "update")
			if err := wf.RunInBackground(updateJobName, cmd); err != nil {
				log.Printf("Error starting update check: %s", err)
			}
		}

		if wf.UpdateAvailable() {
			wf.Configure(aw.SuppressUIDs(true))
			log.Println("Update available!")
			wf.NewItem("An update is available!").
				Subtitle("⇥ or ↩ to install update").
				Valid(false).
				Autocomplete("workflow:update").
				Icon(&aw.Icon{Value: "update-available.png"})
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	wf.Run(func() {
		if err := rootCmd.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	})
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

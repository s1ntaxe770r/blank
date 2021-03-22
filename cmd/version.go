package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versioncmd)
}

var blankver = " you are running blank version 0.1"

var versioncmd = &cobra.Command{
	Use:   "version",
	Short: "prints version of blank",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(color.RedString("ERROR: version command does not require any arguments"))
			os.Exit(1)
		}
		fmt.Println(color.GreenString(blankver))
	},
}

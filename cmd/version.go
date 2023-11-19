package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the parallel world's code that Minaris is at.",
	Long:  `Print the parallel world's code that Minaris is at.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("World line code: v0.1.0")
	},
}

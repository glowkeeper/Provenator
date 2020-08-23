package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version 0.1.0",
	Long: `provenator version 0.1.0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provenator version 0.1.0 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

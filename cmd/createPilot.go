package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var comCrtCmd = &cobra.Command{
	Use:   "comCrtCmd",
	Short: "it can create a new pilot in the database.",
	Long:  `comCrtCmd this is a long description for this comCrt command.`,
	//Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

	},
}

func init() {
	rootCmd.AddCommand(comCrtCmd)
}

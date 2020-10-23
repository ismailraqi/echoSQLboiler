package cmd

import (
	"github.com/ismailraqi/echoSQLboiler/routers"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start it's a command that's allow you to serve the api",
	Long:  `start this is a long description for this start command.`,
	// Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		routers.StartRouters()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

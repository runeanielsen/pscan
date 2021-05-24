package cmd

import (
	"github.com/spf13/cobra"
)

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Add hosts with the add command",
	Long: `Manages the hosts list for pScan

Add hosts with the add command
Delete hosts with the delete command
List hosts with the list command.`,
}

func init() {
	rootCmd.AddCommand(hostsCmd)
}

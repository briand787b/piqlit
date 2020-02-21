package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "echoes all arguments",
	Long:  `All arguments and flags are echoed.  Useful for debugging`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args: %v", args)
	},
}

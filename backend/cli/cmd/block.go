package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/briand787b/piqlit/core/postgres"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	timeout *time.Duration = pflag.Duration("timeout", 5*time.Second, "timeout duration")
)

func init() {
	rootCmd.AddCommand(blockCmd)
}

var blockCmd = &cobra.Command{
	Use:   "block [options]",
	Short: "Block until all connections are healthy",
	Long:  `Block until all connections are healthy`,
	Run: func(cmd *cobra.Command, args []string) {
		pflag.Parse()
		go func() {
			time.Sleep(*timeout)
			log.Fatalln("timed out")
		}()

		fmt.Println("waiting for Postgres connection...")
		postgres.GetDB()
		fmt.Println("done")
	},
}

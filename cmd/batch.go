/*
Copyright © 2022 Rexford A. Nyarko rexfordnyrk@gmail.com

*/
package cmd

import (
	"github.com/rexfordnyrk/webshot/capture"

	"github.com/spf13/cobra"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch [path-to-list-of-urls.txt]",
	Args: cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Run: takeBatch,
}

func init() {
	rootCmd.AddCommand(batchCmd)
}

func takeBatch(cmd *cobra.Command, args []string)  {
	capture.GetBatchScreenShot(args[0], &conf)
}

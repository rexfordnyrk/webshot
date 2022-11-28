/*
Copyright © 2022 Rexford A. Nyarko rexfordnyrk@gmail.com

*/
package cmd

import (
	"github.com/rexfordnyrk/webshot/capture"
	"github.com/spf13/cobra"
)

// singleCmd represents the single command
var singleCmd = &cobra.Command{
	Use:   "single [URL]",
	Args: cobra.ExactArgs(1),
	Short: "Takes a single screenshot of the provided URL",
	Run: takeSingle,
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
func takeSingle(cmd *cobra.Command, args []string)  {
	capture.GetSingleScreenShot(args[0], &conf)
}

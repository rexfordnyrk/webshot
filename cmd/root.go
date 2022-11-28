/*
Copyright © 2022 Rexford A. Nyarko rexfordnyrk@gmail.com

*/
package cmd

import (
	"github.com/rexfordnyrk/webshot/capture"
	"log"
	"os"

	"github.com/spf13/cobra"
)


var (
	conf = capture.Config{}
)
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webshot",
	Short: "Capture Screenshots of web pages.",
	Long: `A simple CLI program that captures screenshots of web pages.
You can have single URL screenshots or batch with a list of screenshots.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&conf.Size, "size", "s","windowed", "size of screenshot, fullscreen or windowed. the default windowed shot size is (1440*900)")
	rootCmd.PersistentFlags().IntVarP(&conf.Width, "width", "w",1440, "The width for both fullscreen and windowed screenshot.")
	rootCmd.PersistentFlags().IntVarP(&conf.Height, "height", "H",900, "The height for a windowed screenshot. This option is ignored for a fullscreen capture.")
	rootCmd.PersistentFlags().StringVarP(&conf.Format, "format", "f","png", "The format of the file output, values is either 'pdf' or 'png'")

	err := conf.ValidateConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}



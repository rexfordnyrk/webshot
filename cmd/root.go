/*
Copyright © 2022 Rexford A. Nyarko rexfordnyrk@gmail.com

*/
package cmd

import (
	"fmt"
	"github.com/rexfordnyrk/webshot/capture"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// PersistencePreRunE on the root command works well
		return initViperConfig(cmd)
	},
	//can be removed after testing as it is not needed in this program
	Run: func(cmd *cobra.Command, args []string) {
	},
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
	rootCmd.PersistentFlags().StringP("config", "c","webshot.json", "Specify the config file name to load. It must be in a JSON format")
	rootCmd.PersistentFlags().StringP("configDir", "d",".", "Specify the directory to find the config file to load. Default is the current directory")

}

func initViperConfig(cmd *cobra.Command) error {
	//initialize viper
	v := viper.New()

	//specifying the type of configuration file format to be read
	v.SetConfigType("json")

	// name of config file obtained from the --config or -c flag
	v.SetConfigName(cmd.Flags().Lookup("config").Value.String())

	//look for config in the User's Home directory
	v.AddConfigPath(cmd.Flags().Lookup("configDir").Value.String())

	// Find and read the config file
	if err := v.ReadInConfig(); err != nil { // Handle errors reading the config file
		// It doesn't matter if config file does not exist we will later try ENV variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//temporarily log loaded config
	fmt.Println("the settings",v.AllSettings())

	return nil
}
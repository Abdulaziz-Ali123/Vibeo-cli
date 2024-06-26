/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Vibeo-cli/cmd/convert"
	"Vibeo-cli/cmd/transcribe"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Vibeo-cli",
	Short: "A comand line to transcribe audio",
	Long: `A comand line to transcribe audio and covert video/audio files to mp3 
	
For example:

Vibeo-cli convert -f example/directory/file
or
Vibeo-cli transcribe -f example/direcrory/audio.mp3 -o transcribed_file
or
Vibeo-cli transcribe -l https://youtube.com/example-video -o transcribed_file
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalette() {
	rootCmd.AddCommand(transcribe.TranscribeCmd)
	rootCmd.AddCommand(convert.ConvertCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Vibeo-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommandPalette()
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transcribe

import (
	"github.com/spf13/cobra"
)

// TranscribeCmd represents the transcribe command
var TranscribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribes audio give youtube link or file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transcribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transcribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package convert

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts video files to mp3",
	Long: `Converts video files to mp3 
For example:
Vibeo-cli convert -f example/directory/video`,
	Run: func(cmd *cobra.Command, args []string) {
		File, _ := cmd.Flags().GetString("file")

		if File != "None" {
			fmt.Println(File)
		} else {
			cmd.Help()
		}

	},
}

func init() {
	ConvertCmd.Flags().StringP("file", "f", "None", "Path to file desired for conversion")
}

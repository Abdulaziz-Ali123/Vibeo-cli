/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package convert

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts video files to mp3",
	Long: `Converts video files to mp3 
For example:
Vibeo-cli convert -f example/directory/video.mp4 -o audio_file`,
	Run: func(cmd *cobra.Command, args []string) {
		File, _ := cmd.Flags().GetString("file")
		Output, _ := cmd.Flags().GetString("output")

		if File != "None" {
			fmt.Println("Converting file...")
			Output += ".wav"

			exec_cmd := exec.Command("ffmpeg", "-i", File, Output)

			b, err := exec_cmd.CombinedOutput()
			if err != nil {
				log.Printf("Convertion failed %v", err)
			}

			fmt.Printf("%s\n", b)
		} else {
			cmd.Help()
		}

	},
}

func init() {
	ConvertCmd.Flags().StringP("file", "f", "None", "Path to file desired for conversion")
	ConvertCmd.Flags().StringP("output", "o", "audio", "Name of the converted file")
}

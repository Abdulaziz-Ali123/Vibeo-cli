/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transcribe

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// TranscribeCmd represents the transcribe command
var TranscribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribes audio given youtube link or file",
	Long: `Transcribes audio given youtube link or file. For example:

Vibeo-cli transcribe -f example/direcrory/audio.mp3 -o transcribed_file
or
Vibeo-cli transcribe -l https://youtube.com/example-video -o transcribed_file`,
	Run: func(cmd *cobra.Command, args []string) {
		File, _ := cmd.Flags().GetString("file")
		Link, _ := cmd.Flags().GetString("link")
		Output, _ := cmd.Flags().GetString("output")

		if File != "None" {
			//pass file to ai made with python speech_recognition module
			fmt.Println(Output)
		} else if Link != "None" {
			exec_cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", Link, "-o", Output+".mp3")
			fmt.Println("Retriving Audio...")
			b, err := exec_cmd.CombinedOutput()
			if err != nil {
				log.Printf("Could not retrive audio file %v", err)
			}

			fmt.Printf("%s\n", b)

		} else {
			cmd.Help()
		}

	},
}

func init() {
	TranscribeCmd.Flags().StringP("file", "f", "None", "Path to file desired for conversion")
	TranscribeCmd.Flags().StringP("link", "l", "None", "youtube link to file desired for conversion")
	TranscribeCmd.Flags().StringP("output", "o", "Transcribe", "Name of the converted file")
}
